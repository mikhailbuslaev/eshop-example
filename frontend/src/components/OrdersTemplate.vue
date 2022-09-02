<template>
<div id="orders-container">
    <div v-for="order in orders" :key="order.id">
        <div class="orders-item">
            <h3>{{ order.cost }} ₽</h3>
            <h4>{{ order.id }}</h4>
        </div>
    </div>
</div>
</template>

<script>
import axios from "axios";
export default {
    name: 'OrdersTemplate',
    data() {
        return {
            orders: [],
            token: '',
            ordersLimit: 20,
            ordersBefore: 0
        };
    },
    methods: {
        loadOrders(startRow, quantity) {
            var bodyFormData = new FormData();
            bodyFormData.append('token', this.token);
            bodyFormData.append('quantity', quantity);
            bodyFormData.append('startrow', startRow);
            axios({
                method: 'post',
                url: 'http://localhost:1111/api/orders',
                data: bodyFormData
            })
            .then(response => {
                this.orders.push(...response.data.message);
            });
        }
    },
    beforeMount() {
        if (this.$cookies.isKey("token")) {
            this.token = this.$cookies.get("token");
        } else {
            this.$router.push({ path: '/login' });
        }
        this.ordersBefore = this.$route.params.page*this.ordersLimit;
        this.getOrders(this.ordersBefore, this.ordersLimit);
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>