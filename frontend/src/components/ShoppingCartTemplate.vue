<template>
<div id="shoppingcart-body">
<h1>your shopping cart:</h1>
    <div v-for="[cartItemId, cartItem] in cartItems" :key="cartItemId">
          <img class="card-picture" :src="cartItem.cardData.picturepath"/>
          <h3>{{ cartItem.cardData.title }}</h3>
          <h4>{{ cartItem.cardData.price }}*{{ cartItem.itemData.count }} ₽</h4>
    </div>
</div>
</template>

<script>
import axios from "axios";
export default {
    name: 'ShoppingCartTemplate',

    data() {
        return {
        cartItems: new Map(),
        emptyCard: {id:'', title:'', price:0, picturepath:'', count:''},
        userId: 0
        };
    },

    methods: {
    getShoppingCart() {
        var bodyFormData = new FormData();
        bodyFormData.append('id', this.userId);
        axios({
        method: 'post',
        url: 'http://localhost:1111/api/shopping_cart',
        data: bodyFormData
        })
        .then((response) => {
            response.data.message.items.forEach(element => {
                this.cartItems.set(element.id, {itemData:element, cardData:this.emptyCard});
                this.getCard(element.id, element.cardid);
            });
        });
      },
    getCard(cartId = '', cardId = '') {
        var bodyFormData = new FormData();
        bodyFormData.append('id', cardId);
        axios({
        method: 'post',
        url: 'http://localhost:1111/api/card',
        data: bodyFormData
        })
        .then((response) => {
            let bufObj = this.cartItems.get(cartId);
            this.cartItems.set(bufObj.itemData.id, {itemData:bufObj.itemData, cardData:response.data.message});
        });
      }
    },

    beforeMount() {
        this.userId = 1;
        this.getShoppingCart();
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>