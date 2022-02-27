package thinkutils

import "ThinkGOGameServer/thinkutils/logger"

var (
	log *logger.LocalLogger = logger.DefaultLogger()

	DateTime     datetime
	StringUtils  stringutils
	RandUtils    randutils
	MD5Utils     md5utils
	IPUtils      iputils
	UUIDUtils    uuidutils
	ThinkMysql   thinkmysql
	ThinkRedis   thinkredis
	FileUtils    fileutils
	RegularUtils regulartils
	JSONUtils    jsonutils
	KafkaUtils   kafkautils
	UDPUtils     udputils
)
