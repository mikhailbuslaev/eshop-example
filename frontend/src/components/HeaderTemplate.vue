<template>
<div id="header-container">
  <nav id="logo-container">
    <router-link to="/catalog/0">
      <h1>Online shop</h1>
    </router-link>
  </nav>
  <nav id="links-container">
      <router-link to="/user/shopping_cart">Shopping Cart</router-link>
      <router-link to="/user/orders/0">My orders</router-link>
      <router-link to="/about">About</router-link>
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
#header-container {
  height:150px;
  background-color: rgb(200, 200, 200);
  display:flex;
  justify-content:space-between;
}


#logo-container {
  padding: 20px;
}

#links-container {
  padding: 40px;
}

nav a {
  padding: 20px;
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: rgb(200, 100, 100);
}

h1 {
  text-align:center;
  margin:0px;
}

</style>