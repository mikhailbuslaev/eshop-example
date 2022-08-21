<template>
<div id="authform-wrapper">
    <div id="authform-body">
        <form @submit="authRequest">
        <label for="fname">
            <h4>Username:</h4>
        </label>
        <input type="text" v-model="userId">
        <label for="lname">
            <h4>Password:</h4>
        </label>
        <input type="text" v-model="password">
        <input type="submit" value="Login">
        </form>
    </div>
    <div id="autherror-body" v-if="isError">
        <h4>Auth error, try another login/password</h4>
    </div>
</div>
</template>

<script>
import axios from "axios";
export default {
  name: 'AuthorizationForm',
  data() {
    return {
      userId: '',
      password: '',
      isError: false
    };
  },

  methods: {
    authRequest() {
        var bodyFormData = new FormData();
        bodyFormData.append('id', this.userId);
        bodyFormData.append('password', this.password);
        axios({
            method: 'post',
            url: 'http://localhost:1111/api/get_token',
            data: bodyFormData
        })
        .then((response) => {
            this.isError = false;
            this.$cookies.set('token', response.data.message, '30min', '', '', true);
            this.$cookies.set('userId', this.userId, '30min', '', '', true);
            this.$router.push({ path: '/catalog/0' });
        })
        .catch((error) => {
            this.isError = true;
            console.log(error);
        });
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#autherror-body {
    margin: auto;
    color:white;
    width:500px;
    height:100px;
    background-color:red;
    border-radius:10px;
}
</style>
