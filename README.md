# xantha
eXtensible Aware aNd Thinking Helper Application

>> What is XANTHA?
The XANTHA project aims to support work automation with the goal of eliminating unnecessary work.
The main goals of the project are:

1. To realize the needs of human business owners for artificially intelligent and adaptable automated employees.
2. To provide an intuitive interface for human business owners to teach and interact with their automated employees.
3. To foster collaboration in the community through development of knowledge modules for the employee AI systems.

>> Why these goals?
Work automation has multiple benefits.

1. Empowerment of small businesses to scale and save money on HR, PTO, human error costs, office space, etc.
2. Reduction of menial labor, allowing greater leisure time and ability to pursue careers of choice for many
3. As 3D printing made home manufacturing possible, XANTHA and related technologies can make business ownership possible for a single person.

However, greater automation reduces the number of jobs available today. In the future, the benefits may be far greater. But it is up to policymakers to realize the economic effects of this technology.

>> What software projects are part of XANTHA?

x-workers
Each "employee" is called an "x-worker".
X-workers make up the employee base for a given organization. They are also designed to interact with human employees and with each other.
Together, they make up a distributed and automated organization.
X-workers are modular in design, meaning that the supervisor has total control over what the X-worker does and doesn't know.

x-view
X-view is a web based interface for managing all x-workers from a central point.
New interaction techniques for x-view are being considered.
Communication with x-workers, their enabling or disabling, changing their settings, and performing "learning sessions" can all be done through x-view.

>> How does it work?

The knowledge originates in a knowledge base which the x-worker consults.
Using entailment, the x-worker can use deductive logic to reach conclusions of truth about the world from a limited set of facts in order to reach knowledge that was not explicitly programmed. An example - 
Say I tell my x-worker that apples each cost $1 and that our fruit budget is $50.
If I later ask if we can afford 100 apples, the x-worker can deduce the answer to this question from the knowledge that it has.
x-workers are powered by a base set of knowledge from Roget's list of concepts used in his thesaurus.
They learn based on human teaching and reinforcement.
Because of this, they may make errors in logic or judgement, like any human worker. However, they never need a vacation, will not make a human error or error in calculation, have a guaranteed uptime 24/7, and don't need coffee to sustain themselves.

x-view provides an ontology view of what the x-worker knows and has learned. Experts can reach in and tweak this knowledge at will to gain total control over the x-worker.

Modules can be created for external devices such as robots and sensors. The x-workers can be programmed to text, call, or otherwise communicate with you at specified times or when certain events occur. These can be managed in x-view to even create "compound actions" by combining a "lift" and "drop" to create a "throw" action, say for a robot arm.

>> A lot of similar projects exist, what makes XANTHA unique?

Products like Siri and Viv are designed as personal assistants. Their job is to manage things for one person - the user. 
X-workers may have multiple demands and interests and need to optimize their queues and workflows. This project is all about making the best automated employees possible that are generalizable to different needs.
Many, many development and operational automation tools exist today. AI controls customer service and scheduling. However, large scale automation is not yet fully trusted for truly running a business including orders, shipments, decision-making and other vital tasks that cost both time and money.

For the purposes of a future where human labor is not needed and robots "run the world", the singularity is not needed. Not even strong AI is needed. However, we do need to take the responsibility to train programs and give them a working knowledge of what they need to know to reliably make decisions and judgements, since they too are becoming workers in our world. We will document our progress here in this repository to share with the world.

>> Technologies

Go
Python
JavaScript - Backbone.js, Angular.js, React.js
JSON/REST
XLearn - a small language facilitating module creation
