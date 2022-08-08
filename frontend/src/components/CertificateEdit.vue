<template>
  <div class="container">
    <div class="row">
      <form-tag class="row" @formEvent="submitHandler" name="myForm" event="formEvent">
        <div class="form-group">
          <text-input
              v-model.trim="certificate.brand"
              label="Model:"
              type="text"
              name="brand"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <div class="row">
            <div class="col-4">
              <text-input
                  v-model.trim="certificate.type_vehicle"
                  label="Tip:"
                  type="text"
                  name="type_vehicle"
                  :required=true
                  :readonly="readonly"
                  :horizontal=true
                  :styleInput=styleInput
                  :styleLabel=styleLabel>
              </text-input>
            </div>
            <div class="col-4">
              <text-input
                  v-model.trim="certificate.variant"
                  label="Varijanta:"
                  type="text"
                  name="variant"
                  :required=true
                  :readonly="readonly"
                  :horizontal=true
                  :styleInput=styleInput
                  :styleLabel=styleLabel>
              </text-input>
            </div>
            <div class="col-4">
              <text-input
                  v-model.trim="certificate.version_vehicle"
                  label="Verzija:"
                  type="text"
                  name="version_vehicle"
                  :required=true
                  :readonly="readonly"
                  :horizontal=true
                  :styleInput=styleInput
                  :styleLabel=styleLabel>
              </text-input>
            </div>
          </div>

          <text-input
              v-model.trim="certificate.commercial_name"
              label="Komercijalna oznaka:"
              type="text"
              name="commercial_name"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.estimated_production_year"
              label="Procenjena godina proizvodnje:"
              type="number"
              name="estimated_production_year"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.max_mass"
              label="Najveca dozvoljena masa vozila(kg):"
              type="number"
              name="max_mass"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.running_mass"
              label="Masa vozila spremnog za voznju(kg):"
              type="number"
              name="running_mass"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.category"
              label="Kategorija vozila:"
              type="text"
              name="category"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.bodywork_code"
              label="Oznaka oblika za karoseriju:"
              type="text"
              name="bodywork_code"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.axles_tyres_num"
              label="Broj osovina i tockova:"
              type="text"
              name="axles_tyres_num"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.length"
              label="Duzina vozila:"
              type="number"
              name="length"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.width"
              label="Sirina vozila:"
              type="number"
              name="width"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.height"
              label="Visina vozila:"
              type="number"
              name="height"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.tyre_wheel"
              label="Pneumatik/naplatak, kombinacija:"
              type="text"
              name="tyre_wheel"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.engine_code"
              label="Oznaka motora:"
              type="text"
              name="engine_code"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.engine_capacity"
              label="Radna zapremina motora"
              type="number"
              name="engine_capacity"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.engine_power"
              label="Najveca neto snaga motora(kW):"
              type="number"
              name="engine_power"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.fuel"
              label="Pogonsko gorivo:"
              type="text"
              name="fuel"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.power_weight_ratio"
              label="Najveca neto snaga/masa vozila(kW/kg):"
              type="text"
              name="power_weight_ratio"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.seat_number"
              label="Broj mesta za sedenje:"
              type="number"
              name="seat_number"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.standing_number"
              label="Broj mesta za stajanje:"
              type="number"
              name="standing_number"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.coupling_device_approval"
              label="Uredjaj za spajanje vucnog i prikljucnog vozila:"
              type="text"
              name="coupling_device_approval"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.max_speed"
              label="Najveca brzina (za vozila vrste L)(km/h):"
              type="number"
              name="max_speed"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.gas_level"
              label="Nivo izduvne emisije/ CO2(g/km):"
              type="text"
              name="gas_level"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.max_laden_mass_axios"
              label="Najvece dozvoljeno osovinsko opterecenje(kg):"
              type="text"
              name="max_laden_mass_axios"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.number_wvta"
              label="Broj sertifikata homologacije tipa vozila:"
              type="text"
              name="number_wvta"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.pollution_cert"
              label="Broj sertifikata za izduvnu emisiju(datum):"
              type="text"
              name="pollution_cert"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.noise_cert"
              label="Broj sertifikata za obuku(datum):"
              type="text"
              name="noise_cert"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

        </div>
        <hr>
        <div class="col-sm-5">
          <input type="submit" v-if="this.action === 'add'" class="btn btn-primary" style="font-size: 14px" value="Snimi">
          <input type="submit" v-if="this.action === 'update'" class="btn btn-primary" style="font-size: 14px" value="Azuriranje">
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
// import VueSingleSelect from "vue-single-select"

export default {
  name: 'CertificateEdit',
  props: {
    action: {
      default: 'add',
      type: String
    },
    certificateId: {
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
    // fileSelected() {
      // if (this.fuelConsumption.bill_file === '') {
      //   return false;
      // }
      // return true;
    // }
  },

  data() {
    return {
      certificate: {brand: "", type_vehicle: "", variant: "", version_vehicle: "", commercial_name: "", estimated_production_year: 0, max_mass: 0, running_mass: 0, category: "", bodywork_code: "", axles_tyres_num: "", length: 0, width: 0, height: 0, tyre_wheel: "", engine_code: "", engine_capacity: 0, engine_power: 0, fuel: "", power_weight_ratio: "", seat_number: 0, standing_number: 0, max_speed: 0, gas_level: "", max_laden_mass_axios: "", number_wvta: "", pollution_cert: "", noise_cert: "", coupling_device_approval: ""},
      users: [],
      styleInput: "font-size: 10px",
      styleLabel: "font-size: 12px",
    }
  },
  methods: {
    onFileChange() {
      const file = this.$refs.file.files[0];
      const reader = new FileReader()
      reader.onloadend = () => {
        const fileString = reader.result;
        this.fuelConsumption.bill_file = fileString;
      }
      reader.readAsDataURL(file);
    },
    removeFile() {

    },
    getDate(date) {
      return date.split('T')[0];
    },
    async submitHandler() {
      this.certificate.estimated_production_year = parseInt(this.certificate.estimated_production_year);
      this.certificate.max_mass = parseInt(this.certificate.max_mass);
      this.certificate.running_mass = parseInt(this.certificate.running_mass);
      this.certificate.length = parseInt(this.certificate.length);
      this.certificate.width = parseInt(this.certificate.width);
      this.certificate.height = parseInt(this.certificate.height);
      this.certificate.engine_capacity = parseInt(this.certificate.engine_capacity);
      this.certificate.engine_power = parseInt(this.certificate.engine_power);
      this.certificate.seat_number = parseInt(this.certificate.seat_number);
      this.certificate.standing_number = parseInt(this.certificate.standing_number);
      this.certificate.max_speed = parseInt(this.certificate.max_speed);
      if (this.certificateId !== '') {
        await this.updateCertificate();
      } else {
        await this.createCertificate();
      }
    },
    async updateCertificate() {
      await axios.post('/certificate/update', JSON.stringify(this.certificate)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        router.push("/certificates");
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
    async createCertificate() {
      await axios.post('/certificate/create', JSON.stringify(this.certificate)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        router.push("/certificates");
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
    if (this.certificateId !== '') {
      axios.get('/certificate/id/' + this.certificateId).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.certificate = JSON.parse(response.data.Data);
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

<style scoped>

</style>