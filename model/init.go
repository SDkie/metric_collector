package model

func Init() {
	InitMongo()
	InitPg()
	InitRedis()
}
