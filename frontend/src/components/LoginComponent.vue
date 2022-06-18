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
          <input type="button" mode="flat" @click="switchMode" class="btn btn-secondary" value="Registracija">
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
      mode: "login",
    }
  },
  methods: {
    submitHandler() {
      console.log("i am in")
      const payload = {
        email: this.email,
        password: this.password,
      }

      axios.post('/users/login', JSON.stringify(payload)).then(function(response){
          console.log(response.data)
        }).catch(function (error) {
          console.log(error);
        });

        // useStore().state.userId =

        router.push('/')
    },
    },
    switchMode() {
      if (this.mode === 'login') {
        this.mode = 'register'
      }
    }
  }
</script>