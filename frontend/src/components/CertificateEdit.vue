<template>
  <div class="container">
    <div class="row">
      <form-tag class="row" @formEvent="submitHandler" name="myForm" event="formEvent">
        <div class="form-group">
          <div class="row mt-2" style="font-size: 12px">
            <label class="col-2" for="radio1">Dokument:</label>
            <div class="col-1 form-check">
              <input type="radio" class="form-check-input" id="radio1" name="sourceDocType" value="COC" :disabled="readonly" v-model="certificate.source_document_type">COC
              <label class="form-check-label" for="radio1"></label>
            </div>
            <div class="col-1 form-check">
              <input type="radio" class="form-check-input" id="radio2" name="sourceDocType" value="TIF" :disabled="readonly" v-model="certificate.source_document_type">TIF
              <label class="form-check-label" for="radio2"></label>
            </div>
            <div class="col-1 form-check">
              <input type="radio" class="form-check-input" id="radio3" name="sourceDocType" value="DEKRA" :disabled="readonly" v-model="certificate.source_document_type">DEKRA
              <label class="form-check-label"></label>
            </div>
            <div class="col-4 form-check">
              <input type="radio" class="form-check-input" id="radio4" name="sourceDocType" value="Potvrda predstavnika proizvođača" :disabled="readonly" v-model="certificate.source_document_type">Potvrda predstavnika proizvođača
              <label class="form-check-label"></label>
            </div>
          </div>
          <text-input
              v-model.trim="certificate.brand"
              label="Marka:"
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
              type="text"
              name="estimated_production_year"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.max_mass"
              label="Najveća dozvoljena masa vozila(kg):"
              type="text"
              name="max_mass"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.running_mass"
              label="Masa vozila spremnog za vožnju(kg):"
              type="text"
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
              label="Broj osovina i točkova:"
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
              label="Dužina vozila:"
              type="text"
              name="length"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.width"
              label="Širina vozila:"
              type="text"
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
              type="text"
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
              label="Radna zapremina motora:"
              type="text"
              name="engine_capacity"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.engine_power"
              label="Najveća neto snaga motora(kW):"
              type="text"
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
              label="Najveća neto snaga/masa vozila(kW/kg):"
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
              type="text"
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
              type="text"
              name="standing_number"
              :required=true
              :readonly="readonly"
              :horizontal=true
              :styleInput=styleInput
              :styleLabel=styleLabel>
          </text-input>

          <text-input
              v-model.trim="certificate.coupling_device_approval"
              label="Uređaj za spajanje vučnog i priključnog vozila:"
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
              label="Najveća brzina (za vozila vrste L)(km/h):"
              type="text"
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
              label="Najveće dozvoljeno osovinsko opterećenje(kg):"
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

<!--          <input v-if="!fileSelected && !readonly" type="file" ref="file" @change="onFileChange" />-->
<!--          <div v-if="fileSelected && !readonly" class="d-flex flex-row">-->
<!--            <button  class="iconBtn" title="Obriši" @click="removeFile">-->
<!--              <i class="fa fa-remove"></i>-->
<!--            </button>-->
<!--            <p>{{certificate.filename}}</p>-->
<!--          </div>-->

<!--          <div v-if="fileSelected && readonly" class="d-flex flex-row">-->
<!--            <button  class="iconBtn" title="Preuzmi" @click="downloadFile">-->
<!--              <i class="fa fa-download"></i>-->
<!--            </button>-->
<!--            <p>{{certificate.filename}}</p>-->
<!--          </div>-->

        </div>
        <hr>
        <div class="col-sm-5">
          <input type="submit" v-if="this.action === 'add'" class="btn btn-primary" style="font-size: 14px" value="Snimi">
          <input type="submit" v-if="this.action === 'update'" class="btn btn-primary" style="font-size: 14px" value="Ažuriranje">
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
    fileSelected() {
      if (this.certificate.file_content === '') {
        return false;
      }
      return true;
    }
  },

  data() {
    return {
      certificate: {source_document_type: "", brand: "", type_vehicle: "", variant: "", version_vehicle: "", commercial_name: "", estimated_production_year: "", max_mass: "", running_mass: "", category: "", bodywork_code: "", axles_tyres_num: "", length: "", width: "", height: "", tyre_wheel: "", engine_code: "", engine_capacity: "", engine_power: "", fuel: "", power_weight_ratio: "", seat_number: "", standing_number: "", max_speed: "", gas_level: "", max_laden_mass_axios: "", number_wvta: "", pollution_cert: "", noise_cert: "", coupling_device_approval: "", file_content: "", filename: ""},
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
        this.certificate.file_content = fileString;
        this.certificate.filename = file.name;
        console.log(this.certificate.filename);
      }
      reader.readAsDataURL(file);
    },
    removeFile() {
      this.certificate.file_content = "";
      this.certificate.filename = "";
    },
    downloadFile() {
      console.log(this.certificate.file_content);
      const arr = this.certificate.file_content.split(',')
      var sampleArr = this.base64ToArrayBuffer(arr[1]);
      const blob = new Blob([sampleArr])

      const link = document.createElement('a')
      link.href = URL.createObjectURL(blob)
      link.download = this.certificate.filename
      link.click()
      URL.revokeObjectURL(link.href)
    },
    base64ToArrayBuffer(base64) {
      var binaryString = window.atob(base64);
      var binaryLen = binaryString.length;
      var bytes = new Uint8Array(binaryLen);
      for (var i = 0; i < binaryLen; i++) {
        var ascii = binaryString.charCodeAt(i);
        bytes[i] = ascii;
      }
      return bytes;
    },
    getDate(date) {
      return date.split('T')[0];
    },
    async submitHandler() {
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