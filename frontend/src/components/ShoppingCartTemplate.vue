<template>
<div id="shoppingcart-body">
<h1>your shopping cart:</h1>
    <div v-for="[cartItemId, cartItem] in cartItems" :key="cartItemId">
        <div class="shoppingcart-item">
            <img class="card-picture" :src="cartItem.cardData.picturepath"/>
            <h4>{{ cartItem.cardData.title }}</h4>
            <h4>{{ cartItem.cardData.price }}*{{ cartItem.itemData.count }} = {{this.renderItemCost(cartItem.cardData.price, cartItem.itemData.count)}}₽</h4>
            <div id="removeitem-button" >
                <button v-on:click="deleteItem(cartItem.itemData.id)">
                <h4>delete item</h4>
                </button>
            </div>
        </div>
    </div>
<h2>summary: {{ showSummaryCost }}₽</h2>
</div>

</template>

<script>
import axios from "axios";
export default {
    name: 'ShoppingCartTemplate',

    data() {
        return {
        cartId: 0,
        token: '',
        cartItems: new Map(),
        emptyCard: {id:'', title:'', price:0, picturepath:'', count:''},
        userId: 0,
        summaryCost: 0
        };
    },

    methods: {
    getShoppingCart() {
        var bodyFormData = new FormData();
        bodyFormData.append('token', this.token);
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
            let obj = this.cartItems.get(cartId);
            this.increaseSummaryCost(this.renderItemCost(obj.itemData.count, obj.cardData.price));
        });
      },
    renderItemCost(first='1', second='1') {
        return parseInt(first)*parseInt(second);
      },
    increaseSummaryCost(itemCost = 0) {
        this.summaryCost += itemCost;
      },
    calculateSummaryCost() {
        this.summaryCost = 0;
        this.cartItems.forEach((value) => {
            this.increaseSummaryCost(this.renderItemCost(value.itemData.count, value.cardData.price));
        });
      },
    deleteItem(cartItemId = '') {
        var bodyFormData = new FormData();
        bodyFormData.append('token', this.token);
        bodyFormData.append('cartitemid', cartItemId);
        axios({
        method: 'post',
        url: 'http://localhost:1111/api/shopping_cart/delete_item',
        data: bodyFormData
        })
        .then(
            this.cartItems.delete(cartItemId),
            this.calculateSummaryCost()
        );
      }
    },

    beforeMount() {
        this.summaryCost = 0;
        if (this.$cookies.isKey("token") && this.$cookies.isKey("userId")) {
            this.userId = this.$cookies.get("userId");
            this.token = this.$cookies.get("token");
        } else {
            this.$router.push({ path: '/login' });
        }
        this.getShoppingCart();
    },

    computed: {
        showSummaryCost: function() {
            return this.summaryCost
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.shoppingcart-item {
    display: flex;
    flex-direction: row;
    gap: 15%;
    justify-content:center;
}
</style>