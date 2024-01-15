function factorial(i) {
	var ret = 1;
	for (;i>1;i--) {
		ret = ret * i;
	}
	return ret
}

function nChooseR(n,r) {
	return factorial(n)/(factorial(r)*factorial(n-r))
}

var count = 0
for (var i = 22; i<101;i++) {
	for (var j=1; j<i;j++) {
		if (nChooseR(i,j) > 1000000) {
			count++;
		}
	}
}

console.log(count);