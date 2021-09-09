console.log("hello");
var url = 'http://192.168.0.14';
var request = new XMLHttpRequest();
request.responseType = 'text';
request.mode='no-corse'
request.onload = function() {
  console.log(request.response);
  setfreq(request.response);
};
function mr(){
  request.open('GET', url);
  request.send();
}
setInterval(mr,150);


