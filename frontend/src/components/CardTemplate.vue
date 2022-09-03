<template>
<div id="card-container">
  <div id="cardpicture-container">
    <img id="cardpicture" :src="card.picturepath"/>
  </div>
  <div id="cardinfo-container">
    <h3>{{ card.title }}</h3>
    <h4>{{ card.price }} ₽</h4>
    <h5 id="carddescription">{{ card.description }}</h5>

    <div id="item-counter">
      <button class="counter-button" id="minus" v-on:click="decrementItemCount">-</button>
      <h4 id="counter-button-text"> count: {{ showItemCount }} </h4>
      <button class="counter-button" id="plus" v-on:click="incrementItemCount">+</button>
    </div>

    <div id="addtocart-button-container" >
      <button id="addtocart-button" v-on:click="addToCart">
      <h4 id="addtocart-button-text">Add to cart</h4>
      </button>
    </div>
  </div>
</div>
</template>

<script>
import axios from "axios";
export default {
  name: 'CardTemplate',
  data() {
    return {
      card: {},
      itemCount: 1,
      cartId: 1
    };
  },

  methods: {
    getCard() {
    var bodyFormData = new FormData();
    bodyFormData.append('id', this.$route.params.id);
    axios({
      method: 'post',
      url: 'http://localhost:1111/api/card',
      data: bodyFormData
    })
    .then((response) => {
        this.card = response.data.message;
      });
    },

    addToCart() {
    if (this.$cookies.isKey("token") && this.$cookies.isKey("userId")) {
      this.userId = this.$cookies.get("userId");
      this.token = this.$cookies.get("token");
    } else {
      this.$router.push({ path: '/login' });
    }
    
    var bodyFormData = new FormData();
    bodyFormData.append('cardid', this.card.id);
    bodyFormData.append('count', this.itemCount);
    bodyFormData.append('token', this.token);
    axios({
      method: 'post',
      url: 'http://localhost:1111/api/shopping_cart/add_item',
      data: bodyFormData
    })
    .then(
        alert('item successfully added to cart')
      );
    },
    incrementItemCount() {
      this.itemCount++;
    },
    decrementItemCount() {
      if (!(this.cartId===0)) {
        this.itemCount--;
      }
    }
  },

  beforeMount() {
    this.getCard();
  },

  computed: {
    showItemCount: function() {
      return this.itemCount
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#card-container {
  width: 100%;
  display:flex;
  justify-content:space-evenly;
}

#item-counter {
  justify-content:center;
  display:flex;
  wrap:nowrap;
}

#cardpicture {
  width: 500px;
  margin:5%;
}

#addtocart-button-container {
  justify-content:center;
  display:flex;
  wrap:nowrap;
}

#carddescription {
  max-width:500px;
}

.counter-button {
  background-color:white;
  border:0px;
  font-size:25px;
  color:rgb(75, 75, 75);
  width:25px;
  padding:0px;
  border-radius:50%;
}

#plus:hover {
  color:white;
  background-color: green;
}

#minus:hover {
  color: white;
  background-color: red;
}

#addtocart-button {
  background-color:white;
  border:0px;
  border-radius:15px;
  font-size:20px;
  color:rgb(75, 75, 75);
}

#addtocart-button:hover {
  background-color:green;
  color: white;
}

h5 {
  color: rgb(150, 150, 150);
}

#counter-button-text {
  margin:5px;
}

#addtocart-button-text {
  margin:10px;
}
</style>