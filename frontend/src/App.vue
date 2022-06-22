<template>
  <template v-if="!isLoggedIn">
    <LoginComponent></LoginComponent>
  </template>
  <template v-else>
    <HeaderComponent></HeaderComponent>
        <div class="col-sm-2">
          <SideBar></SideBar>
        </div>
        <div class="col-sm-10">
          <router-view></router-view>
        </div>
    <FooterComponent></FooterComponent>
  </template>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
import FooterComponent from "./components/FooterComponent.vue";
import LoginComponent from "@/components/LoginComponent";
import SideBar from "@/components/SideBar";

export default {
  name: 'App',
  components: {
    SideBar,
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
  @import './assets/css/main.css';
</style>
