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
		log.Println(query, human.Name, err)
	}

	return engine.GetHumanIDByIdentity(human.Identity)
}

// HumanWithIDExists 返回数据表中是否存在传入自然人ID的对象。
func (engine *Engine) HumanWithIDExists(id int) bool {
	query := `select count(*) from "Human" where "HumanID"=:1`
	result := engine.db.QueryRow(query, id)

	var count int
	if err := result.Scan(&count); err != nil {
		log.Println(query, id, err)
	}
	return count >= 1
}

// HumanWithIdentityExists 返回数据表中是否存在对应身份证的对象。
func (engine *Engine) HumanWithIdentityExists(identity string) bool {
	query := `select count(*) from "Human" where "Identity"=:1`
	result := engine.db.QueryRow(query, identity)

	var count int
	if err := result.Scan(&count); err != nil {
		log.Println(query, identity, err)
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
		log.Println(query, identity, err)
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
		log.Println(query, human.Name, err)
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
		log.Println(query, id, err)
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
		log.Println(query, id, err)
	}
}

// DeleteHumanByIdentity 删除指定身份证的自然人
func (engine *Engine) DeleteHumanByIdentity(identity string) {
	query := `delete from "Human" where "Identity"=:1`
	_, err := engine.db.Exec(query, identity)
	if err != nil {
		log.Println(query, identity, err)
	}
}

// TestTableHuman 测试表Human
func (engine *Engine) TestTableHuman() {
	log.Println("Testing table Human.")

	var (
		testBirthday     = time.Date(2019, time.May, 28, 0, 0, 0, 0, time.UTC)
		testSex          = "男"
		testIdentity     = "123456789012345678"
		testIdentity2    = "098765432112345678"
		testNotes        = "这是测试使用的字段，若此数据项可见，则数据表测试可能没有成功。"
		testPasswordHash = "One$Two$Three"
	)

	engine.DeleteHumanByIdentity(testIdentity)
	engine.DeleteHumanByIdentity(testIdentity2)

	testHuman := Human{
		Name: "测试人",
	}
	testHuman.Identity = testIdentity

	testHumanID := engine.CreateHuman(testHuman)
	if testHumanID == 0 {
		log.Panicln("Table Human test failed: ID should NOT be empty.")
	}
	if !engine.HumanWithIDExists(testHumanID) {
		log.Panicln("Table Human test failed: HumanWithIDExists should return TRUE.")
	}
	if !engine.HumanWithIdentityExists(testHuman.Identity) {
		log.Panicln("Table Human test failed: HumanWithIdentityExists should return TRUE.")
	}

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

	engine.DeleteHumanByID(testHumanID)
	if engine.HumanWithIDExists(testHumanID) {
		log.Panicln("Table Human test failed: HumanWithIDExists should return FALSE.")
	}
	if engine.HumanWithIdentityExists(testHuman.Identity) {
		log.Panicln("Table Human test failed: HumanWithIdentityExists should return FALSE.")
	}
}
