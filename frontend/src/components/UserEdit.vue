<template>
  <div class="container">
    <div class="row">
      <div class="col-5 mx-auto">
        <h1 class="mt-5">Registracija korisnika</h1>
        <hr>
        <form-tag @formEvent="submitHandler" name="myForm" event="formEvent">
          <text-input
              v-model.trim="name"
              label="Ime"
              type="text"
              name="name"
              required="true">
          </text-input>

          <text-input
              v-model.trim="lastName"
              label="Prezime"
              type="text"
              name="lastName"
              required="true">
          </text-input>

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
          <input type="submit" class="btn btn-primary m-2" value="Registracija">
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
import TextInput from "@/components/forms/TextInput";
import FormTag from "@/components/forms/FormTag";
import axios from "axios";
import router from "@/router";

export default {
  name: 'UserEdit',
  components: {FormTag, TextInput},
  data() {
    return {
      name: "",
      lastName: "",
      email: "",
      password: "",
    }
  },
  methods: {
    async submitHandler() {
      const payload = {
        first_name: this.name,
        last_name: this.lastName,
        email: this.email,
        password: this.password,
      }

      await axios.post('/users/register', JSON.stringify(payload)).then((response) => {
        if (response.data.Data === "") {
          this.errorMsg = "Error during login!";
          console.log("Token is empty: " + response.data.Data);
          return;
        }
      }, (error) => {
        console.log(error);
      });

      await router.push("/users");
    },
  }
}
</script>