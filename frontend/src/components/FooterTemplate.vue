<template>
<div id="footer-container">
  <div class="footer-element">
    <a href="https://www.w3schools.com"><h3>Link example</h3></a>
    <a href="https://www.w3schools.com"><h3>Another link example</h3></a>
    <a href="https://www.w3schools.com"><h3>Some link</h3></a>
  </div>
  <div class="footer-element">
    <a href="https://www.w3schools.com"><h3>Link</h3></a>
    <a href="https://www.w3schools.com"><h3>Link to another place here</h3></a>
  </div>
  <div class="footer-element">
    <a href="https://www.w3schools.com"><h3>Footer link</h3></a>
    <a href="https://www.w3schools.com"><h3>Random link</h3></a>
  </div>
  <div class="footer-element">
    <a href="https://www.w3schools.com"><h3>Link example</h3></a>
    <a href="https://www.w3schools.com"><h3>Link</h3></a>
  </div>
</div>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#footer-container {
  display:flex;
  justify-content:space-around;
  height:200px;
  background-color: #ccc;
}

.footer-element {
  width:250px;
  padding:20px;
}

h3 {
  text-align:left;
}

a {
  color:rgb(75, 75, 75);
  text-decoration: none;
}
</style>