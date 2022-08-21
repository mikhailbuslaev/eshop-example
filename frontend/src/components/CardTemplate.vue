<template>
<h1>card view example</h1>
<div id="card-body">
    <img class="card-picture" :src="card.picturepath"/>
    <h3>{{ card.title }}</h3>
    <h4>{{ card.price }} ₽</h4>
    <h5>{{ card.description }}</h5>

    <div id="item-counter">
      <button v-on:click="decrementItemCount">-</button>
      <h4>item count:{{ showItemCount }}</h4>
      <button v-on:click="incrementItemCount">+</button>
    </div>

    <div id="addtocart-button" >
      <button v-on:click="addToCart">
      <h4>add to cart</h4>
      </button>
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
#card-body {
  width: 500px;
  background: #ccc;
  margin: 1em auto;
}

#item-counter {
  justify-content:center;
  display:flex;
  wrap:nowrap;
}

.card-picture {
  width: 500px;
}

#addtocart-button {
  justify-content:center;
  display:flex;
  wrap:nowrap;
}
</style>