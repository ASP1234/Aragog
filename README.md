# Aragog [![Build Status](https://travis-ci.com/ASP1234/Aragog.svg?branch=master)](https://travis-ci.com/ASP1234/Aragog)
![Meet Aragog](assets/Aragog.png?raw=true)
Aragog is an Internet spider bot that systematically browses the World Wide Web, typically for the purpose of web scraping so he can feed himself and his family. 

## Traits
* **Fast** - He has a huge family and so with their help he manages to get things done quickly.
* **Hard Working** - He and his family retries a certain number of times if they encounter a failure before putting their hopes down.
* **Helpful** - He and his family always strives to help their master by informing how things are going and if there is something panicky!!!
* **Flexible** - He has a progessive mind and so he allows his master to change the way he and his family fetches, reproduce, communicate or store their food.
* **Obedient** - He and his family obeys rules and restrictions feeded to him.

## How his family works as a whole?
![Alt text](assets/OverallArchitecture.png?raw=true "Title")
* He receives a seed from Master.
* He fetches the seed and store it into food storage.
* Each seed can have links to other sources for food, so he process those links and push them over his family communication medium.
* His family member looks for any message and if it's there he/she go to hunt the food and repeat the lifecycle till they are all satisfied.
* Once satisfied, Aragog informs the master and returns to his home.

## How he and his family process the message?
![Alt text](assets/ProcessorArchitecture.png?raw=true "Title")
* Active member fetches the food.
* He/She then puts it into their *shared* food storage.
* Master has taught them how to filter the child links. So, the he/she uses all the configured skills to filter the child links that needs to be fetched.
  
  Example:
  
  * Member removes links that his/her family has already visited within the configured time. Because, he.she knows that they have fetched it already and within the configured time it is unlikely that there will be something new to fetch.
  * Member removes links that does not belong to the domain they are targetting. Else, they might get lost in the huge world.
  * Member removes fake links because they don't serve anyone good.
  
  **NOTE:**
  These skills are configurable and master can choose to use whatever skill he/she wants or perhaps develop a new skill altogether.
  
 * He/She then publishes the filtered links over their family communication medium. As the process is tiresome, so he/she spawns children to process the filtered links and returns home.
 
   **NOTE:**
   Even though the members spawn children, they don't want to cause over population and hence they make sure that at any point of time their family members count doesn't cross the configured threshold.
   
 ## How they store food?
 As of now, they are storing it within their limited home. But, if needed they can shift to a cloud store as they are highly flexible. Like everyone else, Aragog also doesn't know what the future looks like. So, he stores with **NoSQL** mindset to meet the **everchanging info** his family may need to store and the sheer **volume** of it.
 
 Currently, they persist the following fields:
 
 address | lastModificationDate | links | retryAttempts | status
 --------| -------------------- | ----- | ------------- | ------
 
 ## What's next?
 *"They believed that I was the monster that dwells in what they call the Chamber of Secrets."* - Aragog
 
 Aragog and his family are obedient to his master but they are getting complaints of privacy invasion. So, they are looking to improve their skills by understanding the *robots.txt* file and use it as a guiding principle in their invasions.
 
 Aragog wishes survival of his family and so he is willing to shift to *cloud* and distribute senior members among their *multiple homes*.
