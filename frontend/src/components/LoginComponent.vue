<template>
  <div class="container">
    <div class="row">
      <div class="col-5 mx-auto">
        <h1 class="mt-5">Prijava</h1>
        <hr>
        <form-tag @formEvent="submitHandler" name="myForm" event="formEvent">
          <text-input
              v-model.trim="email"
            label="Email"
            type="email"
            name="email"
            required="true">
          </text-input>

          <text-input
              v-model.trim="password"
              label="Sifra"
              type="password"
              name="password"
              required="true">
          </text-input>
          <hr>
          <input type="submit" class="btn btn-primary m-2" value="Prijava">
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
import TextInput from "@/components/forms/TextInput";
import FormTag from "@/components/forms/FormTag";
import axios from "axios";
// import {useStore} from 'vuex';
import router from './../router/index.js'

export default {
  name: 'LoginComponent',
  components: {FormTag, TextInput},
  data() {
    return {
      email: "",
      password: "",
      errorMsg: "",
      token: "",
    }
  },
  methods: {
    async submitHandler() {
      const payload = {
        email: this.email,
        password: this.password,
      }

      await axios.post('/users/login', JSON.stringify(payload)).then((response) => {
        if (response.data.Data === "") {
          this.errorMsg = "Error during login!";
          console.log("Token is empty: " + response.data.Data);
          return;
        }
        this.token = response.data.Data;
        }, (error) => {
        console.log(error);
      });

      console.log("Sve ok", this.token);

       this.$store.dispatch('setTokenAction', this.token);

       await router.push("/");
    },
    }
  }
</script>