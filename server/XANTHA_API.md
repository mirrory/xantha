GET POST PUT DELETE /organization
GET POST PUT DELETE /user
GET POST PUT DELETE /xworker
GET POST PUT DELETE /role
GET POST PUT DELETE /usermessage
GET POST PUT DELETE /job
GET POST PUT DELETE /task
GET POST PUT DELETE /action

Design v.1
Organizations(id,name,number_of_xworkers)
UsersInOrgs(user_id,org_id)
Users(id,username,first_name,last_name,title,image,security_group)
XWorkersInOrgs(xworker_id,org_id)
XWorkers(id,name,title,image,permission_group)
XWorkerMemory(id,this,that,it)
XWorkerLongMemory(id,entity)
XWorkerRoles(xworker_id,role_id)
RolesInOrgs(role_id,org_id)
Roles(id,name)
ActionsInRole(action_id,role_id)
UserMessages(id,sender,recipient,content)
XWorkerMessages(id,sender,recipient,content)
Jobs(id,xworker_id,date_assigned,date_updated,date_scheduled,is_recurring,next_scheduled_run,date_completed,is_past,is_in_progress,percent_done)
Tasks(id,job_id,task_type,sequence_number)
ActionsInTask(action_id,task_id)
Actions(id,action_type,permission_needed)

Flow:
User creates account and Organization.
User creates XWorker for Organization.
User assigns Role to XWorker.
User sends UserMessage to XWorker.
XWorker replies with XWorkerMessage.
User can set up Job with Tasks manually. 
User can set up Actions manually.
Otherwise, XWorker determines what Tasks and Jobs based on UserMessages.

cron pkg

Lex - 
Separate sentences.
Remove punctuation and capitalization.

Construct a CST.

Parse - 
Recursive tail-end parse.

Construct an AST.

Semantic Analysis - 
Words are tied to Modules.
Match word to module. 
Make sure can access the previous and subsequent token.
Push entire sentence to relevant Modules.

Noun - an Instance or Type
Verb - an Action
Adjective - a Property
Adverb - a notation on the Action, perhaps a Time or Place
Logical Connectives

Determine sentence type - Declarative, Interrogative, Imperative, Exclamatory

Have randomized synonymous responses for greater immersion

Return response string and links to created jobs and/or actions taken.

Back-End components:
api/ - serves the api
nlp/ - parses language and produces responses in nl
thoughtspace/ - accepts a state and produces job(s)
concepts/ - thesaurus modules
dbLayer/ - communicates with the database
cron/ - handles scheduling

Message string is received by API at /usermessage
String goes into NLP to produce a State(s)
States go into thoughtspace to become Jobs
thoughtspace uses relevant concepts and previous data from database
cron schedules jobs / thunks to be run at a later time
thoughtspace indicates what jobs were created, updated, or destroyed
nlp forumulates response based on this
response string goes out through api

message -> state -> job -> cron -> response

Example:
"Buy 2 apples."
State: {verb: "buy", qty: 2, object: "apples"}
Concepts: Transaction, Quantity, Food
Transaction: Do we have enough money?
Quantity: Are there enough already?
Food: Where do we get them? What budget do we use?
Job output: "purchase apples 2"
Cron "now" is implied - Run the Job. Send purchase request to supplier
If don't know supplier, should ask User "What supplier do you want me to use?"
Or even "What type of apples?"
These should be stored in memory for later - 
Can change x-worker mode to either ask a lot of questions each time, or -
Just do what you were told last time, and ask less questions.
Depends on what kind of employee you want.
Finally, explain that job completed successfully or needs more info
NLP: "Done! I bought 2 apples from Amazon Fresh."
Ta-da!