package dbcd

import (
	"log"
	"time"
)

// Human 是表“Human”的模型
type Human struct {
	HumanID      int
	Name         string
	Sex          string
	Birthday     time.Time
	Identity     string
	Notes        string
	PasswordHash string
}

// CreateHuman 在数据表中创建Human数据项，并返回创建Human数据项的HumanID。
func (engine *Engine) CreateHuman(human Human) int {
	query := `insert into "Human" ("Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash")
values (:1, :2, :3, :4, :5, :6)`

	_, err := engine.db.Exec(query, human.Name, human.Sex, human.Birthday, human.Identity, human.Notes, human.PasswordHash)
	if err != nil {
		Trace(err, query, human.Name)
	}

	return engine.GetHumanIDByIdentity(human.Identity)
}

// ExistHumanWithID 返回数据表中是否存在传入自然人ID的对象。
func (engine *Engine) ExistHumanWithID(id int) bool {
	query := `select count(*) from "Human" where "HumanID"=:1`
	result := engine.db.QueryRow(query, id)

	var count int
	if err := result.Scan(&count); err != nil {
		Trace(err, query, id)
	}
	return count >= 1
}

// ExistHumanWithIdentity 返回数据表中是否存在对应身份证的对象。
func (engine *Engine) ExistHumanWithIdentity(identity string) bool {
	query := `select count(*) from "Human" where "Identity"=:1`
	result := engine.db.QueryRow(query, identity)

	var count int
	if err := result.Scan(&count); err != nil {
		Trace(err, query, identity)
	}
	return count >= 1
}

// GetHumanIDByIdentity 返回身份证对应自然人的ID。
// 若身份证不存在于数据表中，则返回0。
func (engine *Engine) GetHumanIDByIdentity(identity string) int {
	query := `select "HumanID" from "Human" where "Identity"=:1`
	result := engine.db.QueryRow(query, identity)

	var id int
	if err := result.Scan(&id); err != nil {
		Trace(err, query, identity)
		return 0
	}
	return id
}

// GetHumanByID 返回指定ID的自然人。
func (engine *Engine) GetHumanByID(id int) *Human {
	query := `select "HumanID", "Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash" from "Human" where "HumanID"=:1`
	result := engine.db.QueryRow(query, id)
	var human Human
	if err := result.Scan(&human.HumanID, &human.Name, &human.Sex, &human.Birthday, &human.Identity, &human.Notes, &human.PasswordHash); err != nil {
		Trace(err, query, human.Name)
		return nil
	}
	return &human
}

// GetHumanByIdentity 返回指定身份证的自然人。
func (engine *Engine) GetHumanByIdentity(identity string) *Human {
	query := `select "HumanID", "Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash" from "Human" where "Identity"=:1`
	result := engine.db.QueryRow(query, identity)
	var human Human
	if err := result.Scan(&human.HumanID, &human.Name, &human.Sex, &human.Birthday, &human.Identity, &human.Notes, &human.PasswordHash); err != nil {
		return nil
	}
	return &human
}

// UpdateHumanByID 更新指定ID的Human。
func (engine *Engine) UpdateHumanByID(human Human, id int) {
	query := `update "Human" set "Name"=:1, "Sex"=:2, "Birthday"=:3, "Identity"=:4, "Notes"=:5, "PasswordHash"=:6 where "HumanID"=:7`
	_, err := engine.db.Exec(query, human.Name, human.Sex, human.Birthday, human.Identity, human.Notes, human.PasswordHash, id)
	if err != nil {
		Trace(err, query, id)
	}
}

// UpdateHumanByIdentity 更新指定身份证的Human。
func (engine *Engine) UpdateHumanByIdentity(human Human, identity string) {
	id := engine.GetHumanByIdentity(identity).HumanID
	engine.UpdateHumanByID(human, id)
}

// DeleteHumanByID 删除指定id的自然人。
func (engine *Engine) DeleteHumanByID(id int) {
	query := `delete from "Human" where "HumanID"=:1`
	_, err := engine.db.Exec(query, id)
	if err != nil {
		Trace(query, id, err)
	}
}

// DeleteHumanByIdentity 删除指定身份证的自然人
func (engine *Engine) DeleteHumanByIdentity(identity string) {
	query := `delete from "Human" where "Identity"=:1`
	_, err := engine.db.Exec(query, identity)
	if err != nil {
		Trace(err, query, identity)
	}
}

// TestTableHuman 测试表Human
func (engine *Engine) TestTableHuman() {
	log.Println("Testing table Human.")

	// 准备测试环境。
	var (
		testBirthday     = time.Date(2019, time.May, 28, 0, 0, 0, 0, time.UTC)
		testSex          = "男"
		testIdentity     = "123456789012345678"
		testIdentity2    = "098765432112345678"
		testNotes        = "（测试字段）"
		testPasswordHash = "（测试密码哈希）"
	)
	testHuman := Human{
		Name: "（测试自然人）",
	}
	testHuman.Identity = testIdentity
	engine.DeleteHumanByIdentity(testIdentity)
	engine.DeleteHumanByIdentity(testIdentity2)

	// 测试CREATE和READ。
	testHumanID := engine.CreateHuman(testHuman)
	if testHumanID == 0 {
		log.Panicln("Table Human test failed: ID should NOT be empty.")
	}
	if !engine.ExistHumanWithID(testHumanID) {
		log.Panicln("Table Human test failed: HumanWithIDExists should return TRUE.")
	}
	if !engine.ExistHumanWithIdentity(testHuman.Identity) {
		log.Panicln("Table Human test failed: HumanWithIdentityExists should return TRUE.")
	}

	// 测试UPDATE。
	testHuman.Birthday = testBirthday
	testHuman.Sex = testSex
	testHuman.Identity = testIdentity2
	testHuman.Notes = testNotes
	testHuman.PasswordHash = testPasswordHash
	engine.UpdateHumanByIdentity(testHuman, testIdentity)

	testHuman2 := engine.GetHumanByIdentity(testIdentity2)
	if testHuman2 == nil {
		log.Panicln("Table Human test failed: testHuman2 should not be nil.")
	}
	if !testHuman2.Birthday.Equal(testBirthday) {
		log.Panicln("Table Human test failed: testHuman2.Birthday", testHuman2.Birthday, "is not consist with testBirthday", testBirthday, ".")
	}
	if testHuman2.Sex != testSex {
		log.Panicln("Table Human test failed: testHuman2.Sex is not consist with testSex.")
	}
	if testHuman2.Notes != testNotes {
		log.Panicln("Table Human test failed: testHuman2.Notes is not consist with testNotes.")
	}

	// 测试DELETE。
	engine.DeleteHumanByID(testHumanID)
	if engine.ExistHumanWithID(testHumanID) {
		log.Panicln("Table Human test failed: HumanWithIDExists should return FALSE.")
	}
	if engine.ExistHumanWithIdentity(testHuman.Identity) {
		log.Panicln("Table Human test failed: HumanWithIdentityExists should return FALSE.")
	}
}
