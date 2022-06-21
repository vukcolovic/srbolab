<template>
  <template v-if="!isLoggedIn">
    <LoginComponent></LoginComponent>
  </template>
  <template v-else>
    <HeaderComponent></HeaderComponent>
    <router-view></router-view>
    <FooterComponent></FooterComponent>
  </template>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
import FooterComponent from "./components/FooterComponent.vue";
import LoginComponent from "@/components/LoginComponent";
// import {useStore} from 'vuex';
// import {computed} from 'vue';

export default {
  name: 'App',
  components: {
    LoginComponent,
    HeaderComponent,
    FooterComponent,
  },
  computed : {
    isLoggedIn : function(){
      console.log("Auth", this.$store.getters.isAuthenticated);
      return this.$store.getters.isAuthenticated;
    }
  },
  beforeCreate() {
    this.$store.commit('initialiseStore');
  }
}
</script>

<style>
</style>
