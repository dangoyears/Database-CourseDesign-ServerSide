# 运行此脚本将导入大量测试数据

import random
import requests
from public import URL, POST_DATA as token


# 随机姓名
names = ['汲冬雪', '霜昊英', '夏侯长霞', '安宵雨', '南宫清妙', '风云韶', '俞米琪', '表北', '仍忆雪', '丘慧雅', '赏萧曼', '历安歌', 
'以嘉怡', '之寄琴', '郜元柳', '瓮宏阔', '凭甜', '劳丽泽', '戈依霜', '天晶滢', '昂慧婕', '鄞紫南', '撒暄妍', '令若雁', '汝平彤',
'阳玲', '拜娟秀', '万俟淑哲', '马高峻', '泷岑', '商嘉瑞', '苍楠楠', '逄夜南', '笪望', '隆醉蝶', '始运升', '钟离以柳', '妫思怡', 
'泰和正', '系书雁', '端木炫明', '骑叶彤', '赛恨瑶', '归安柏', '典星鹏', '琴鹍', '戎友安', '性鑫', '綦玲琅', '翦心宜', 
'段雪瑶', '辉语蕊', '申珠星', '费听露', '茅淳静', '玄衍', '弘冠玉', '鲜于语儿', '璩量', '嘉曾琪', '仇琼音', '袁恬静', '祝夏容',
'令狐和风', '宇天薇', '修暮芸', '承烁', '其永', '巧蕴美', '虎香卉', '翠凝阳', '是觅荷', '寒友安', '姓鹏', '湛奇水', '丛弘博', 
'市学名', '检昆琦', '务绍晖', '函晴丽', '阴夏兰', '萧笑', '益恨寒', '厉秀慧', '孙诗双', '南门丹云', '夔慈', '裴朗然', 
'甄朵', '拓跋高朗', '东方睿德', '闾丘晓星', '酒逸明', '行致萱', '班听云', '麦雨旋', '穆慧英', '轩辕源', '佟诗怀', '谈书白', '潜丽玉', 
'侯轩', '卢幻竹', '仲嘉勋', '少嘉慕', '杞彦杉', '百承颜', '柏平乐', '门雪晴', '乾悠逸', '冯芳洲', '利亦旋', '禄乐邦', 
'寇冰安', '奈密思', '卞炫明', '钦素昕', '丹源', '从翠阳', '孟彬彬', '尤海瑶', '姬经', '坚骏桀', '雷拔', '藩宏盛', '次流惠', 
'刚兴怀', '贺熙柔', '仙鸿德', '戏燕珺', '双晓莉', '悉秀媛', '太史易云', '莫萌阳', '富察茜茜', '敛嘉丽', '蔚才捷', '焦依琴', 
'错姝艳', '老良', '沃成和', '谢浩瀚', '濮小萍', '肖晓星', '泉香莲', '晏香菱', '何阳兰', '锐彩妍', '毛薇', '尹珂', '匡菱华', 
'粟如云', '国盼旋', '蒯承嗣', '考瑞芝', '告傲菡', '禹瑞灵', '昔德泽', '楼文君', '浑天蓝', '史笑容', '雍景福', '喜雅美', '南静逸',
'度忻', '薛玟玉', '宁浦和', '苌欣美', '澄奇文', '晋天空', '员清淑', '壤驷精', '宾逸春', '仆翔', '库觅山', '合雪兰', '冀访风',
'中德本', '仲孙秀越', '硕春冬', '税觅柔', '斐欣可', '愚青文', '军海', '平英哲', '原运良', '程小春', '桑新知', '纳喇晶滢', 
'歧晓枫', '冼德惠', '农幻儿', '淡妃', '波秋柔', '代水之', '芮伟泽', '能夏旋', '巴嘉悦', '支正豪', '池丹琴', '闭睿识', 
'卑茹云', '奕俊楚', '和初兰', '涂蔚然', '衣博', '台高兴', '斯鹏海', '盘雨双', '秘海菡', '汤巧蕊', '敖凡雁', '厍叶飞', '镇寒天', 
'开含双', '京鹏运', '弥黛娥', '称冰薇', '林念霜', '方嘉庆', '戚秋寒', '揭新竹', '叶静枫', '多凡白', '裘水蓝', '闽翱', '宝听筠', 
'籍语薇', '靳舒云', '莘飞沉', '谭春冬', '靖骞泽', '首悠素', '司寇文', '丑代容', '府代卉', '黄奇思', '后甜恬', '毕然', '闾会雯', 
'旷锐达', '驹虹星', '宏渊', '羽和歌', '营秀丽', '家荣轩', '业灵秀', '都雅媚', '塞初翠', '壬浩博', '岳紫文', '苦刚', '说海雪', 
'完易绿', '花子瑜', '纵星汉', '杨又青', '终初兰', '褒鸿信', '富睿思', '银尔蓉', '聂良畴', '迟修文', '皮同和', '理清婉', '世思淼', 
'謇芷容', '公良耘豪', '睦颖颖', '鲍伟毅', '董开济', '稽水风', '千盼秋', '释梦丝', '答智杰', '亓南霜', '拱晶滢', '汪宛秋', '庹欣可', 
'由瑾瑜', '蓝兴修', '张廖东', '武雅琴', '茹梓美', '洛依秋', '堂凝丝', '谌令慧', '士半烟', '乌孙梅', '谬谷芹', '甘骊娟', '革鸣', 
'穰兰蕙', '奉维运', '郏瑾瑶', '成君', '机采萱', '訾雪卉', '毓鹤梦', '折凝洁', '阿韶华', '钮莹华', '鲁蕙', '仪子薇', '第五兴朝', 
'章佳如彤', '秦莺莺', '车子', '侨巧夏', '偶梓瑶', '柴泽', '蒋彭湃', '卯紫杉', '郁向晨', '生悦怡', '法依波', '鹿曦', '晁千叶', 
'金云蔚', '汗夏柳', '狄凯乐', '贾婉然', '乌诗双', '郑涵涵', '无萧', '淳于曼蔓', '乘乃', '兴卿', '保诗蕾', '蒲兰梦', '宛秋柳', 
'野靖易', '哀代天', '尉觅珍', '塔思烟', '庞建章', '宫丹彤', '陀采文', '舜越泽', '易来', '律曜栋', '介哲彦', '东门为', '幸益', 
'梁丘华彩', '城梦雨', '却涵涤', '温嘉颖', '简明凝', '锁淑君', '荀飞荷', '贰优扬', '慈丽容', '桐中', '牵妞', '虞晓蕾', '普云逸', 
'巫之卉', '兰乐怡', '皋广君', '卷津', '僧悦可', '贲迎蓉', '植冰夏', '葛珹', '斋凯泽', '巢智敏', '隐悦畅', '海以冬', '明冬']

# 创建班级
data = {
    **token, **{
        "college": "计算机科学与网络学院",
        "specialty": "软件工程",
        "grade": "17",
        "class": "1",
    }
}
for i in range(1, 5):
    data["class"] = str(i)
    requests.post(URL + '/write/college', json=data)

# 创建教师
data = {
    **token, **{
        "college": "计算机科学与网络学院",
        "name": "南宫问天",
        "jobId": "1000000001",
        "sex": "男",
        "education": "硕士",
        "graduation": "广州大学",
        "birthday": "2000-1-1",
        "age": "34",
        "idCard": "440582199708310612",
        "password": "dangoyears",
        "position": "教务办主任"
    }
}
requests.post(URL + '/write/teacher', json=data)

# 创建学生
data={
    **token, **{
        "college": "计算机科学与网络学院",
        "specialty": "软件工程",
        "grade": "17",
        "class": "1",
        "name": "学生姓名",
        "studentId": "1706300000",
        "status": "在读本科生",
        "sex": "女",
        "birthday": "1998-9-22",
        "idCard": "440582199708310610",
        "password": "dangoyears",
        "yearSystem": "4"
    }
}

for i in range(1, 5):
    for j in range(1, 40):
        data["class"] = str(i)
        data["name"] = random.choice(names)
        data["sex"] = random.choice(["男", "女"])
        data["birthday"] = "1998-"+ str(random.choice(range(1, 13))) + "-" + str(random.choice(range(1, 29)))
        data["studentId"] = str(1706300001 + i*42+j)
        data["idCard"] = str(random.randint(440582199708310610, 990582199708310610))
        requests.post(URL + '/write/student', json=data)
