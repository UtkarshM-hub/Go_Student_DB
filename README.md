# Go_Student_DB
It is command line tool to manage student database.The tool provides various options like Add, Delete, Update and List students. The data of the students is stored in a file in local machine.

## #Installation Guide
  * Clone this repo into your local directory using following command
  ```bash
  git clone git@github.com:UtkarshM-hub/Go_Student_DB.git
  ```
  * Run following command
  ```bash
  go mod download && go build .
  ```
### Congratulations🥳🎊 you've completed the setup of GO_Student_DB on your local machine

## #Commands

### 1. student
  * This is the root command of the command line tool
  * It displays the list of commands and the example of help command
  ![Screenshot (119)](https://user-images.githubusercontent.com/70505181/192771769-cc63b42b-93ea-4f8b-aef3-ac762a86fe49.png)
  
### 2. Help 
  * Gives information about the command
  * For Example ```student help AddCmd``` gives information about Add command
  ![up4](https://user-images.githubusercontent.com/70505181/192785854-943571b9-b7fe-4b1b-9e89-a05a5039bcf6.png)

### 3. Other commands
  * This contains all the reamining commands like ```Add```, ```Del```, ```Update``` and ```List```
  * How to use each command is exaplanind in the following
    * Add command- ```student Add```
    * List command- ```student List```
    * Update command- ```student Update <INDEX_NUMBER>```
    * Del command- ```student Del <INDEX_NUMBER>```
  #### NOTE- You can see index number of student by using list command
![up3](https://user-images.githubusercontent.com/70505181/192788251-ea5af2a0-71b9-43ca-9be4-a673c81998a1.png)

