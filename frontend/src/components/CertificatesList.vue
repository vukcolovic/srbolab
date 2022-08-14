<template>
    <div class="container">
      <div class="row m-2">
        <div class="btn-group">
          <button class="iconBtn" title="Dodaj" @click="$router.push({name: 'CertificateEdit', params: {certificateId: '', action: 'add' }})">
            <i class="fa fa-plus"></i>
          </button>
          <button class="iconBtn" title="Pregledaj" :disabled="selectedCertificate == null" @click="$router.push({name: 'CertificateEdit', params: {certificateId: selectedCertificate.id, action: 'view' }})" >
            <i class="fa fa-eye"></i>
          </button>
          <button class="iconBtn" title="Izmeni" :disabled="selectedCertificate == null" @click="$router.push({name: 'CertificateEdit', params: {certificateId: selectedCertificate.id, action: 'update' }})" >
            <i class="fa fa-pencil">
            </i>
          </button>
          <button class="iconBtn" title="Obrisi" :disabled="selectedCertificate == null" @click="deleteCertificate">
            <i class="fa fa-trash-o">
            </i>
          </button>
          <button class="iconBtn ms-auto" title="PDF" :disabled="selectedCertificate == null" @click="getPdf()">
             <i class="fa fa-file-pdf-o">
             </i>
          </button>
        </div>
      </div>
          <div class="collapse multi-collapse border" style="font-size: 0.7em" id="filter">
            <div class="row">
            </div>
      </div>
      <div class="row mt-2">
        <vue-table-lite
            @row-clicked="selectCertificate"
            :columns="columns"
            :rows="rows"
            :total= "totalCount"
            @do-search="doSearch"
            :is-loading="isLoading"
        ></vue-table-lite>
      </div>
    </div>
  </template>

  <script>
    import VueTableLite from "vue3-table-lite";
    import axios from "axios";
    import notie from 'notie';

    export default {
      name: 'CertificatesList',
      components: { VueTableLite },
      data() {
        return {
          columns: [
            {
              label: 'ID',
              field: 'id',
              width: '2%',
              isKey: true,
            },
            {
              label: 'Tip',
              field: 'type_vehicle',
              width: '6%',
            },
            {
              label: 'Varijanta',
              field: 'variant',
              width: '10%',
            },
            {
              label: 'Verzija',
              field: 'version_vehicle',
              width: '16%',
            },
            {
              label: 'Masa voz. sprem. vožnju',
              field: 'running_mass',
              width: '12%',
            },
            {
              label: 'Oznaka motora',
              field: 'engine_code',
              width: '10%',
            },
            {
              label: 'Zapremina motora',
              field: 'engine_capacity',
              width: '12%',
            },
            {
              label: 'Najveća neto snaga motora',
              field: 'engine_power',
              width: '12%',
            },
            {
              label: 'Kategorija',
              field: 'category',
              width: '12%',
            },
            {
              label: 'Godina. proiz.',
              field: 'estimated_production_year',
              width: '12%',
            },
          ],
          rows: [],
          selectedCertificate: null,
          selectedRowId: '',
          users: [],
          isLoading: false,
          filterObject: {type_vehicle: '', variant: '', version_vehicle: '', running_mass: '', estimated_production_year: '', engine_code: '', engine_capacity: '', engine_power: '', fuel: '', category: '' },
          totalCount: 0,
          win: ""
        }
      },
      methods: {
        selectCertificate(rowData) {
          this.selectedCertificate = rowData;
        },
        getDate(date) {
          return date.split('T')[0];
        },
        async countCertificates() {
          await axios.post('/certificate/count', JSON.stringify(this.filterObject)).then((response) => {
            if (response.data === null || response.data.Status === 'error') {
              notie.alert({
                type: 'error',
                text: 'Greska: ' + response.data.ErrorMessage,
                position: 'bottom',
              })
              return;
            }
            this.totalCount = response.data.Data;
          }, (error) => {
            notie.alert({
              type: 'error',
              text: "Greska: " + error,
              position: 'bottom',
            })
          });
        },
        async doSearch(offset, limit) {
          this.isLoading = true;
          await axios.post('/certificate/list?skip=' + offset + '&take=' + limit, JSON.stringify(this.filterObject)).then((response) => {
            if (response.data === null || response.data.Status === 'error') {
              notie.alert({
                type: 'error',
                text: 'Greska: ' + response.data.ErrorMessage,
                position: 'bottom',
              })
              return;
            }
            this.rows = JSON.parse(response.data.Data);
          }, (error) => {
            notie.alert({
              type: 'error',
              text: "Greska: " + error,
              position: 'bottom',
            })
          });

          await this.countCertificates();

          this.isLoading = false;
        },
        async deleteCertificate() {
          if (confirm("Da li ste sigurni da zelite da obrisete sertifikat?")) {
            await axios.get('/certificate/delete/' + this.selectedCertificate.id).then((response) => {
              if (response.data === null || response.data.Status === 'error') {
                notie.alert({
                  type: 'error',
                  text: 'Greska: ' + response.data.ErrorMessage,
                  position: 'bottom',
                })
              }
            }, (error) => {
              notie.alert({
                type: 'error',
                text: "Greska: " + error,
                position: 'bottom',
              })
            });
          }

          location.reload();
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
        getPdf () {
            this.win = prompt("Unesite win oznaku vozila:");
            if (this.win.length < 10 || this.win.length > 17) {
              notie.alert({
                type: 'error',
                text: "Duzina WIN oznake ne moze biti ispod 10 i preko 17 karaktera!",
                position: 'bottom',
              })
              return;
            }
            if (this.win == null || this.win == "") {
              return;
            }
          axios.get('/certificate/pdf/id/' + this.selectedCertificate.id + "/win/" + this.win)
              .then(response => {
                var fileContent = JSON.parse(response.data.Data);
                var sampleArr = this.base64ToArrayBuffer(fileContent);
                const blob = new Blob([sampleArr], { type: 'application/pdf' })

                const link = document.createElement('a')
                link.href = URL.createObjectURL(blob)
                link.download = "certificate_" + this.selectedCertificate.id
                link.click()
                URL.revokeObjectURL(link.href)
                //FIXME add notie
              }).catch(console.error)
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
    }
      },
    created() {
        this.doSearch(0, 10);
      },
    mounted() {
        this.getAllUsers();

        var self = this;
        let cols = document.getElementsByClassName("vtl-thead-th");
        for (let i = 1; i < cols.length; i++) {
          const node = document.createElement("input");
          switch (i) {
            case 1:
              node.setAttribute("id", "type_vehicle");
              node.addEventListener('keyup', function (event){
                self.filterObject['type_vehicle'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            case 2:
              node.setAttribute("id", "variant")
              node.addEventListener('keyup', function (event){
                self.filterObject['variant'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            case 3:
              node.setAttribute("id", "version_vehicle")
              node.addEventListener('keyup', function (event){
                self.filterObject['version_vehicle'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            case 4:
              node.setAttribute("id", "running_mass")
              node.addEventListener('keyup', function (event){
                self.filterObject['running_mass'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            case 5:
              node.setAttribute("id", "engine_code")
              node.addEventListener('keyup', function (event){
                self.filterObject['engine_code'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            case 6:
              node.setAttribute("id", "engine_capacity")
              node.addEventListener('keyup', function (event){
                self.filterObject['engine_capacity'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            case 7:
              node.setAttribute("id", "engine_power")
              node.addEventListener('keyup', function (event){
                self.filterObject['engine_power'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            case 8:
              node.setAttribute("id", "category")
              node.addEventListener('keyup', function (event){
                self.filterObject['category'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            case 9:
              node.setAttribute("id", "estimated_production_year")
              node.addEventListener('keyup', function (event){
                self.filterObject['estimated_production_year'] = event.target.value;
                self.doSearch(0, 10)
              })
              break;
            default:
          }
          node.style.width = "100%";
          cols[i].appendChild(node)
        }
      }
    }
  </script>

  <style scoped>
    ::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
      font-size: 10px;
    }
    ::v-deep(.vtl-table td),
    ::v-deep(.vtl-table tr) {
      font-size: 10px;
    }
    ::v-deep(.vtl-table tr):active {
      font-size: 20px;
    }
    ::v-deep(.vtl-paging-count-label),
    ::v-deep(.vtl-paging-page-label),
    ::v-deep(.vtl-paging-count-dropdown),
    ::v-deep(.vtl-paging-page-dropdown){
      font-size: 10px;
    }
    ::v-deep(.vtl-paging-pagination-page-link) {
      font-size: 10px;
    }
  </style>
