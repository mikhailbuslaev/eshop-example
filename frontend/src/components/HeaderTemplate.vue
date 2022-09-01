<template>
<div class="header">
  <h1>Online shop</h1>
  <nav>
      <router-link to="/catalog/0">Home</router-link> | 
      <router-link to="/about">About</router-link> | 
      <router-link to="/user/shopping_cart">Shopping Cart</router-link>
  </nav>
</div>
</template>

<script>
export default {
    name: 'HeaderTemplate'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.header {
  height:200px;
  background-color: rgb(200, 200, 200);
}

h1 {
  text-align:center;
  padding-top:30px;
  margin:0px;
}

nav {
  padding: 40px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: rgb(200, 100, 100);
}
</style>