# xantha
## eXtensible Aware aNd Thinking Helper Application

> What is XANTHA?

The XANTHA project aims to explore the possibilities of generic task automation, and represents my personal feasibility study into human/computer interaction for automated processes specifically.

The main premises of this project are:

1. That human business owners can make use of artificially intelligent and adaptable automated employees.
2. An intuitive interface can allow human business owners to teach and interact with their automated employees.
3. To increase the efficiency of these automated employees, we can develop specialized knowledge modules for the employee AI systems.

> Why are these premises important to explore?

Automation of work is a highly debated topic as of the time of writing. I am personally interested in the social sciences field of "future studies", particularly the human/computer dimension of automation. The vision I imagine is not one where computers replace human workers, but where humans and computers can work together. Computers will take care of menial tasks and drudgery, giving humans more time for creative, fun and meaningful work. 

A few of the benefits include:

1. Empowerment of small businesses to scale and save money on HR, PTO, human error costs, office space, etc.
2. Reduction of menial labor, allowing more time for creative work and leisure time
3. As 3D printing made home manufacturing possible, XANTHA and related technologies can make business ownership possible for a single person

This future becomes more feasible once programs are able to reliably and autonomously make decisions based on state and context. Many business owners do not yet feel comfortable leaving time-critical decisions to computers without human input. XANTHA's model will be to allow humans to "train" automated employees to do their jobs correctly, and the employees will output as much information as possible in a clear format to account for their work history, current tasks, and future scheduled tasks.

> What software projects are part of XANTHA?

### x-workers

Each "employee" is called an "x-worker".

X-workers make up the employee base for a given organization. They are also designed to interact with human employees and with each other.

Together, they make up a distributed and automated organization.

X-workers are modular in design, meaning that the supervisor has total control over what the X-worker does and doesn't know.

In other words, the accountant x-worker does not know how to make Facebook posts in the manner that the social media management x-worker might. Thus, there is no need to worry that your private accounting details might be posted to your business' Facebook. One of the biggest problems with automation platforms today is unpredictability. Transactions are sometimes not atomic and might have side effects, including bugs that can cause unexpected behavior. X-workers will have logs, audit trails and modularization built into their design, lowering the chances of risky errors.

### x-view
X-view is a web based interface for managing all x-workers from a central point.

New interaction techniques for x-view are being considered. Human operators will be able to tweak what x-workers know via visual representations of learned skills. This should allow business logic to be mapped accurately to the actions that x-workers will take in certain situations.

Communication with x-workers, their enabling or disabling, changing their settings, and performing "learning sessions" can all be done through x-view.

> How does it work?

The knowledge originates in a knowledge base which the x-worker consults.

Using entailment, the x-worker can use deductive logic to reach conclusions of truth about the world from a limited set of facts in order to reach knowledge that was not explicitly programmed. 

An example - 

Say I tell my x-worker that apples each cost $1 and that our fruit budget is $50.

If I later ask if we can afford 100 apples, the x-worker can deduce the answer to this question from the knowledge that it has. It can even reply that we'd need 50 more dollars applied to our fruit budget in order to afford 100 apples.

x-workers are powered by a base set of knowledge from Roget's list of concepts used in his thesaurus. The reason for this choice is that the vast majority of concepts, constructs and logical connectives that we use on a daily basis are encoded in our language. These concepts are things we use to make decisions - we may decide to go outside and jog based on the weather, the time of day, our energy level, and multiple other variables. Since these conditions can be described through language, using a compendium of concepts has the potential for a rich knowledge schema.

X-workers also learn based on human teaching and reinforcement.

Because of this, they may make errors in logic or judgement, like any human worker. However, they never need a vacation, will not make a human error or error in calculation, have a guaranteed uptime 24/7 (theoretically), and don't need coffee to sustain themselves.

X-view provides an ontology view of what the x-worker knows and has learned. Experts can reach in and tweak this knowledge at will to gain total control over the x-worker.

Modules can be created for external devices such as robots and sensors. The x-workers can be programmed to text, call, or otherwise communicate with you at specified times or when certain events occur. These can be managed in x-view to even create "compound actions" by combining a "lift" and "drop" to create a "throw" action, say for a robot arm. The possibilties for extensibility are endless.

> A lot of similar projects exist, what makes XANTHA unique?

Products like Siri and Viv are designed as personal assistants. Their job is to manage things for one person - the user. 

X-workers may have multiple demands and interests and need to optimize their queues and workflows. This project is all about making the best automated employees possible that are generalizable to different needs / generic tasks. (The scope obviously seems to be huge. As my personal project to explore, it is fully experimental.)

Many, many development and operational automation tools exist today. AI controls customer service and scheduling. However, large scale automation is not yet fully trusted for truly running a business including orders, shipments, decision-making and other vital tasks that cost both time and money to maintain. And mistakes can cost exponentially more.

For the purposes of a future where human labor is not needed and robots "run the world", the singularity is not needed. Not even strong AI is needed. However, we do need to take the responsibility to train programs and give them a working knowledge of what they need to know to reliably make decisions and judgements, since they too are becoming workers in our world. This project just aims to investigate a few theoretical methods for making automated employees easier to interact with. 

> Technologies

* Go - Back-end
* JavaScript - Front-end - React.js
* JSON/REST - API
* PostgreSQL - Database
* XLearn - a small language facilitating external module creation
* React Native - Mobile app
* Docker - Deployment