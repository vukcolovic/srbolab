<template>
  <div class="container">
    <div class="row">
      <h3 v-if="action === 'add'" class="mt-2">Dodaj</h3>
      <h3 v-if="action === 'view'" class="mt-2">Pregled</h3>
      <h3 v-if="action === 'update'" class="mt-2">Azuriranje</h3>
      <hr>
      <form-tag class="row" @formEvent="submitHandler" name="myForm" event="formEvent">
        <div class="col-sm-5">
          <text-input
              v-model.trim="fuelConsumption.date_consumption"
              label="Datum"
              type="date"
              name="date"
              required="true"
              :readonly="readonly">
          </text-input>

          <text-input
              v-model.trim="fuelConsumption.car_registration"
              label="Registracija vozila"
              type="text"
              name="carRegistration"
              required="true"
              :readonly="readonly">
          </text-input>

          <label class="mb-2">Sipao/la</label>
          <vue-single-select
              name="poured_by"
              @input="(selected) => fuelConsumption.poured_by = selected"
              :value="fuelConsumption.poured_by"
              :options="users"
              option-label="first_name"
          ></vue-single-select>
          <input type="submit" v-if="this.action === 'add'" class="btn btn-primary m-2" value="Dodavanje">
          <input type="submit" v-if="this.action === 'update'" class="btn btn-primary m-2" value="Azuriranje">
        </div>
        <div class="col-sm-5">
          <text-input
              v-model.trim="fuelConsumption.liter"
              label="Kolicina (litar)"
              type="number"
              name="liter"
              :readonly="readonly">
          </text-input>

          <text-input
              v-model.trim="fuelConsumption.price"
              label="Cena (RSD)"
              type="number"
              name="price"
              :readonly="readonly">
          </text-input>

          <label class="mb-2">Vrsta goriva</label>
          <vue-single-select
              name="fuel_type"
              @input="(selected) => fuelConsumption.fuel_type = selected"
              :value="fuelConsumption.fuel_type"
              :options="fuelType"
              option-label=""
          ></vue-single-select>
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
import VueSingleSelect from "vue-single-select"

export default {
  name: 'FuelConsumptionEdit',
  props: {
    action: {
      default: 'add',
      type: String
    },
    fuelConsumptionId: {
      default: '',
      type: String
    }
  },
  components: {FormTag, TextInput, VueSingleSelect},
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
      fuelConsumption: {date_consumption: "", liter: 0.0, price: 0.0, car_registration: "", poured_by: null},
      users: [],
      fuelType: ["BENZIN", "DIZEL", "GAS"],
    }
  },
  methods: {
    async submitHandler() {
      if (this.fuelConsumptionId !== '') {
        await this.updateFuelConsumption();
      } else {
        await this.createFuelConsumption();
      }
    },
    async updateFuelConsumption() {
      await axios.post('/fuel-consumption/update', JSON.stringify(this.fuelConsumption)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        router.push("/fuel");
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
    async createFuelConsumption() {
      await axios.post('/fuel-consumption/create', JSON.stringify(this.fuelConsumption)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        router.push("/fuel");
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
    this.getAllUsers();
    if (this.fuelConsumptionId !== '') {
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