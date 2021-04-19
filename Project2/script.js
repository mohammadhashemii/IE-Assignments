
/*get response from the api*/
function submit() {
	let url = "https://api.genderize.io/?name="
	let name = document.getElementById("name").value;
	
	/*check if the input text is empty and handle the error*/
	if(name.length < 1) {
		window.alert("The Name field is empty!");
	}
	else {
		let resp;
		fetch(url + name)
		.then((response) => handle_HTTP_errors(response))
		.then((response) => response.json())
		.then((data) => set_result(data))
		.then((data) => save_answer(data))
		.then((data) => check_history(data));
	}



}

/*clear the history fot specific data*/
function clear_history() {
	localStorage.removeItem(document.getElementById("name").value);
	document.getElementById("history").style.visibility="hidden";
}

/*check if the data has been stored in the local storage or not*/
function check_history(data) {
	let obj = JSON.parse(JSON.stringify(data));
	let history = localStorage.getItem(obj.name);

	if (history != null){
		document.getElementById("history").style.visibility = "visible";
		document.getElementById("saved-value").innerHTML = history;

	}
	else{
		document.getElementById("history").style.visibility="hidden";
	}
}

/*store the prediction into the local storage*/
function save_answer(data) {
	let rbs = document.querySelectorAll('input[name="gender"]');
	let selectedValue = null;
	
	/*a loop for selecting the radio button options*/
	for (let rb of rbs){
		if (rb.checked){
			selectedValue = rb.value;
		}
	}
	/*check if the user select one radio option or not*/
	if (selectedValue == null){
		return data;
	}

	let obj = JSON.parse(JSON.stringify(data));
	/*insert one record into the local storage*/
	localStorage.setItem(obj.name, selectedValue);

	/*un-check the radio bautton options*/
	document.getElementById("male").checked = false;
	document.getElementById("female").checked = false;

	return data;

}

/*handle HTTP errors*/
function handle_HTTP_errors(response) {
	if (response.status >= 300){
		document.getElementById("error-text").innerHTML = "HTTP error: " + response.statusText;
		document.getElementById("error-text").style.visibility = "visible";
	}
	else {
		document.getElementById("error-text").style.visibility = "hidden";
	} 
	return response;
}

/*set the result of the prediction into the webpage*/
function set_result(data) {
	/*parse the json data to retrieve its items*/
	let obj = JSON.parse(JSON.stringify(data));
	/*check if the api has a prediction fot the input name*/
	if (obj.gender == null){

		document.getElementById("gender").innerHTML = "Gender";
		document.getElementById("probability").innerHTML = "Probability";
		document.getElementById("error-text").innerHTML = "The name is unpredictable!";
		document.getElementById("error-text").style.visibility = "visible";

		return data
	}

	document.getElementById("error-text").style.visibility = "hidden";
	document.getElementById("gender").innerHTML = obj.gender;
	document.getElementById("probability").innerHTML = obj.probability;

	return data;
}