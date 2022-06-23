local installation
1. create your own mysql database
2. run this query
=============================================
CREATE TABLE IF NOT EXISTS `specialteams` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(70) DEFAULT NULL,
  `email` varchar(70) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `address` varchar(120) DEFAULT NULL,
  `image` varchar(150) DEFAULT NULL
  PRIMARY KEY (`id`)
)
=============================================
3. run this command
```bash
  composer install
```
4. setup your .env file
5. run this command to run locally
```bash
  go run main.go
```
6. Happy Coding Dude!