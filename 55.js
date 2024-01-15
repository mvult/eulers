//Receive a number
function isPalindrome(i) {
	var tmp = "" + i
	for (var j = 0; j < tmp.length/2;j++) {
		if (tmp.charAt(j) != tmp.charAt(tmp.length-j-1)) {
			return false
		}
	}
	return true
}

function reverse(i) {
	var tmp = "" + i
	var splitString = tmp.split("")	
	var reverseArr = splitString.reverse()
	return reverseArr.join("")
}

function checkIfLychrel(i) {
	for (var j = 0;j < 51;j++) {
		console.log(i + "   " + parseInt(reverse(i))) 
		if (isPalindrome(i + parseInt(reverse(i)))) {
			console.log(i + parseInt(reverse(i)))
			return true
		}
		i = i + parseInt(reverse(i))
		console.log(i)
	}
	return false
}

console.log(checkIfLychrel())

/*
var count = 0
for (var i = 1;i<10001;i++) {
	if (checkIfLychrel(i)) {
		count++
	}
}
console.log(count);

*/