# gocx
Dialogflow CX agent handy commands. 

# Commands
gocx help
gocx agent get intents $agent-id -dir ${dir} -file ${filename} -verbose
gocx agent get flows $agent-id  -dir ${dir} -file ${filename} ${filename} -verbose
gocx agent get pages $agent-id -dir ${dir} -file ${filename} ${filename} -verbose 
gocx agent server -port=9000