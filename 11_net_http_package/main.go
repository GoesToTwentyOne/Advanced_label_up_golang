package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)

}
func home(response http.ResponseWriter, req *http.Request) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<title>CSS Template</title>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
	#header,#footer{
		background: red;
		padding: 20px;
		text-align: center;
		font-family: fantasy;
		font-display: swap;
		font-size: 25px;
		#body{
			font-family: Verdana,sans-serif;
			font-size: 0.9em;   
		}
	}
	#menu ul{
		list-style: none;
		padding: 0px;
	}
	#menu ul li{
		display: inline;
		margin: 50px;
	}
	#content{
		margin: 5px;
		padding: 10px;
		background-color: lightblue;
	}
	.article{
		margin: 5px;
	  padding: 10px;
	  background-color: white;
	}
	</style>
	</head>
    <body id="body">
        <div id="header">
            <h1>This Protom Alo</h1>
        </div>
        <div id="menu">
            <ul>
                <li>Home</li>
                <li>Life style</li>
                <li>International</li>
            </ul>
        </div>
        <div id="content">
            <h1>news name</h1>
            <div class="article">
                <h1>The Padma</h1>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque in porta lorem. Morbi condimentum est nibh, et consectetur tortor feugiat at.</p>
                
            </div>
            <div class="article">
                <h1>The jamuna</h1>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque in porta lorem. Morbi condimentum est nibh, et consectetur tortor feugiat at.</p>
            </div>
        </div>
        <div id="footer">
            <h1>This is Nihad</h1>
        </div>
       
    </body>
	</html>
	`

	fmt.Fprintf(response, body)
}
func contact(response http.ResponseWriter, req *http.Request) {
	body := `<!DOCTYPE html>
	<html>
	<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
	body {font-family: Arial, Helvetica, sans-serif;}
	* {box-sizing: border-box;}
	
	input[type=text], select, textarea {
	  width: 100%;
	  padding: 12px;
	  border: 1px solid #ccc;
	  border-radius: 4px;
	  box-sizing: border-box;
	  margin-top: 6px;
	  margin-bottom: 16px;
	  resize: vertical;
	}
	
	input[type=submit] {
	  background-color: #4CAF50;
	  color: white;
	  padding: 12px 20px;
	  border: none;
	  border-radius: 4px;
	  cursor: pointer;
	}
	
	input[type=submit]:hover {
	  background-color: #45a049;
	}
	
	.container {
	  border-radius: 5px;
	  background-color: #f2f2f2;
	  padding: 20px;
	}
	</style>
	</head>
	<body>
	
	<h3>Contact Form</h3>
	
	<div class="container">
	  <form action="/action_page.php">
		<label for="fname">First Name</label>
		<input type="text" id="fname" name="firstname" placeholder="Your name..">
	
		<label for="lname">Last Name</label>
		<input type="text" id="lname" name="lastname" placeholder="Your last name..">
	
		<label for="country">Country</label>
		<select id="country" name="country">
		  <option value="australia">Australia</option>
		  <option value="canada">Canada</option>
		  <option value="usa">USA</option>
		</select>
	
		<label for="subject">Subject</label>
		<textarea id="subject" name="subject" placeholder="Write something.." style="height:200px"></textarea>
	
		<input type="submit" value="Submit">
	  </form>
	</div>
	
	</body>
	</html>
	
	`
	fmt.Fprintf(response, body)

}
