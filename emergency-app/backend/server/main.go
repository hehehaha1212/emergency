package main

import (
"log" //used for logs
"os" //used for working with env and systemlevel thingis
"github.com/gin-gonic/gin" //this handles ports and https work
"github.com/joho/godotenv" //this handels  env variable loading
"github.com/username/emergency-connect/config" //imports the configs for postgres
"github.com/username/emergency-connect/internal/routes" //file to manage api
)

func main()
{
 err:= godotenv.load.(".env")
 if err != nil {
   log.fatel("error loading env")
}

config.ConnectDB()
gin.SetMode(gin.ReleaseMode)
router := gin.default()
routes.setuproutes(router)

port  := os.getenv("PORT")
if port == ""{  
   port="8080"
}
log.println("starting server port:"port)
router.Run(":"+port)
}

