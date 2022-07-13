<template>
  <div class="container">
    <div class="row">
      <h3 v-if="action === 'add'" class="mt-2">Dodaj nepravilnost</h3>
      <h3 v-if="action === 'view'" class="mt-2">Pregled</h3>
      <h3 v-if="action === 'update'" class="mt-2">Azuriranje</h3>
      <hr>
      <form-tag class="row" @formEvent="submitHandler" name="myForm" event="formEvent">
      <div class="col-sm-5">
          <text-input
              v-model.trim="irregularity.subject"
              label="Izvestaj"
              type="text"
              name="subject"
              required="true"
              :readonly="readonly">
          </text-input>

          <text-area-input
              v-model.trim="irregularity.description"
              label="Opis"
              type="text"
              name="description"
              required="true"
              :readonly="readonly">
          </text-area-input>

          <label class="mb-2">Ispitivac</label>
          <vue-single-select
              name="Nivo"
              @input="(selected) => irregularity.controller = selected"
              :value="irregularity.controller"
              :options="users"
              option-label="first_name"
          ></vue-single-select>
        <input type="submit" v-if="this.action === 'add'" class="btn btn-primary m-2" value="Dodavanje">
        <input type="submit" v-if="this.action === 'update'" class="btn btn-primary m-2" value="Azuriranje">
      </div>
      <div class="col-sm-5">
        <label class="mb-2">Nivo</label>
        <vue-single-select
            name="level"
            @input="(selected) => irregularity.irregularity_level = selected"
            :value="irregularity.irregularity_level"
            :options="irregularityLevels"
            option-label="code"
        ></vue-single-select>

        <text-area-input
            v-model.trim="irregularity.notice"
            label="Napomena"
            type="text"
            name="notice"
            :readonly="readonly">
        </text-area-input>
      </div>
      </form-tag>
    </div>
  </div>
</template>

<script>
import TextInput from "@/components/forms/TextInput";
import FormTag from "@/components/forms/FormTag";
import axios from "axios";
import router from "@/router";
import notie from 'notie';
import TextAreaInput from "@/components/forms/TextAreaInput";
import VueSingleSelect from "vue-single-select"

export default {
  name: 'IrregularityEdit',
  props: {
    action: {
      default: 'add',
      type: String
    },
    irregularityId: {
      default: '',
      type: String
    }
  },
  components: {FormTag, TextInput, TextAreaInput, VueSingleSelect},
  computed: {
    readonly() {
      if (this.action === 'view') {
        return true;
      }
      return false;
    }
  },
  data() {
    return {
      irregularity: {subject: "", description: "", notice: "", irregularity_level: null, controller: null},
      irregularityLevels: [],
      users: [],
    }
  },
  methods: {
    async submitHandler() {
      if (this.irregularityId !== '') {
        await this.updateIrregularity();
      } else {
        await this.createIrregularity();
      }
    },
    async updateIrregularity() {
      await axios.post('/irregularity/update', JSON.stringify(this.irregularity)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        router.push("/irregularities");
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
    async createIrregularity() {
      await axios.post('/irregularity/create', JSON.stringify(this.irregularity)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        router.push("/irregularities");
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
    async getAllIrregularityLevels() {
      await axios.get('/enumeration/irregularity-levels/all').then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.irregularityLevels = JSON.parse(response.data.Data);
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
    async getAllUsers() {
      await axios.get('/users/list?skip=0&take=100').then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.users = JSON.parse(response.data.Data);
        this.users.forEach(user => user.first_name = user.first_name + ' ' +  user.last_name);
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
     this.getAllIrregularityLevels();
     this.getAllUsers();
    if (this.irregularityId !== '') {
      axios.get('/irregularity/id/' + this.irregularityId).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.irregularity = JSON.parse(response.data.Data);
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