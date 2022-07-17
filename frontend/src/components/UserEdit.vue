<template>
  <div class="container">
    <div class="row">
      <div class="col-sm-5 mx-auto">
        <h3 v-if="action === 'add'" class="mt-2">Registracija</h3>
        <h3 v-if="action === 'view'" class="mt-2">Pregled</h3>
        <h3 v-if="action === 'update'" class="mt-2">Azuriranje</h3>
        <hr>
        <form-tag @formEvent="submitHandler" name="myForm" event="formEvent">
          <text-input
              v-model.trim="user.first_name"
              label="Ime"
              type="text"
              name="name"
              required="true"
              :readonly="readonly">
          </text-input>

          <text-input
              v-model.trim="user.last_name"
              label="Prezime"
              type="text"
              name="lastName"
              required="true"
              :readonly="readonly">
          </text-input>

          <text-input
              v-model.trim="user.email"
              label="Email"
              type="email"
              name="email"
              required="true"
              :readonly="readonly">
          </text-input>

          <text-input
              v-model.trim="user.current_password"
              label="Trenutna sifra"
              type="password"
              name="currentPassword"
              :hidden="!showCurentPassword"
              :required="false"
              :readonly="readonly">
          </text-input>

          <text-input
              :hidden="readonly"
              v-model.trim="user.password"
              :label="passwordLabel"
              type="password"
              name="password"
              :required="passwordRequired"
              :readonly="readonly">
          </text-input>
          <hr>
          <input type="submit" v-if="this.action === 'add'" class="btn btn-primary m-2" value="Registracija">
          <input type="submit" v-if="this.action === 'update'" class="btn btn-primary m-2" value="Azuriranje">
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
import notie from 'notie';

export default {
  name: 'UserEdit',
  props: {
    action: {
      default: 'add',
      type: String
    },
    userId: {
      default: '',
      type: String
    }
  },
  components: {FormTag, TextInput},
  computed: {
    readonly() {
      if (this.action === 'view') {
        return true;
      }
      return false;
    },
    passwordRequired() {
      if (this.action === 'add') {
        return true;
      }
      return false;
    },
    showCurentPassword() {
      if (this.action === 'update') {
        return true;
      }
      return false;
    },
    passwordLabel() {
      if (this.action === 'update') {
        return "Nova sifra";
      }
      return "Sifra";
    }
  },
  data() {
    return {
      user: {first_name: "", last_name: "", email: "", password: "", current_password: ""}
    }
  },
  methods: {
    async submitHandler() {
      if (this.userId !== '') {
        await this.updateUser();
      } else {
        await this.createUser();
      }
    },
    async createUser() {
      await axios.post('/users/register', JSON.stringify(this.user)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        router.push("/users");
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
    async updateUser() {
      await axios.post('/users/update', JSON.stringify(this.user)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        router.push("/users");
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
  },
  mounted() {
    if (this.userId !== '') {
      axios.get('/users/id/' + this.userId).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.user = JSON.parse(response.data.Data);
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    }
  }
}
</script>