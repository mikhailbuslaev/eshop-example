<template>
<div id="authform-wrapper">
    <div id="authform-container">
        <form @submit="authRequest">
        <label for="fname">
            <h3>Username</h3>
        </label>
        <input type="text" v-model="userId">
        <label for="lname">
            <h3>Password</h3>
        </label>
        <input type="text" v-model="password">
        <input type="submit" value="Login">
        </form>
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
            this.$router.go(-1);
        })
        .catch(
            alert('wrong login/password')
        );
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

#authform-container {
    width:100%;
    display:flex;
    justify-content:center;
}

form {
    width:50%;
    display:flex;
    flex-direction:column;   
}

input[type=text] {
    width: 100%;
    padding: 12px 20px;
    margin: 8px 0;
    box-sizing: border-box;
    border-radius:5px;
    border: 1px solid rgb(150, 150, 150);
    font-size: 20px;
}

input[type=submit] {
    width: 100%;
    padding: 12px 20px;
    margin: 8px 0;
    box-sizing: border-box;
    border-radius:5px;
    border: 1px solid rgb(150, 150, 150);
    font-size: 20px;
    color: rgb(75, 75, 75);
}
</style>
