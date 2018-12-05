Installation and User’s Manual 
 
Installation:
 
1. Install docker CE version
 
2.Download HOLMS folder, enter /HOLMS
 
3.Open terminal and run the following command:
 
$ docker-compose up -d --build
 
(In this yaml file, we deploy three containers: go-web-app, node-web-app, mongodb exposed on ports: 8000:8000, 8080:8080, and 27017:27017 respectively. If there are port conflicts on your computer, please $docker stop all container before running docker-compose.yaml file.)
 
 
Visit HOLMS website
 
1.Open Chrome browser, and type http://localhost:8080/HOLMS_main_page.html in the  address bar to visit HOLMS website
 

Main Function
 
1. Sign-Up:
The system will assume you as new to HOLMS, users can try to sign up with clicking sign-up green button showing on screen

2. Log-In:
After sign-up, you will be redirected to sign-in page. Please sign in and enjoy the journey of visiting HOLMS

3.Log-Out:
When you have logged in, you can always log out at side bar.

4. Add Course:
You may try to add course to account. Clicking Adding course button will redirect you to course selection page. Course records are showing on course selection page if you already have added some courses.


5. Taking Courses:
In dashboard page, you may click “go to course” button on the course card to take the courses you have selected. Then you can click “lecture” to watch the corresponding  lecture video. You can also jump to other chapters by clicking chapter buttons on the right.  

6.Assignment Submission:
In course page, you can click “assignment” to complete an assignment by entering and submitting your answer in the text area. 

7. Discussion Board:
You can enter the discussion board and view others’ discussions by through the sidebar. Then you can navigate to the bottom of the discussion board page to post a new discussion. 

8.View & Update Account Information :
You can click “my account” to view and update your account information, including change your password.
 


 

