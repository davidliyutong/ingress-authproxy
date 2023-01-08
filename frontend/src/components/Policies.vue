<template>
  <div>

    <v-data-table
        :headers="headers"
        :items="policyItems"
        sort-by="calories"
        :items-per-page=20
        class="elevation-1"
        no-data-text="No Users"
        show-select
        :search="search"
    >
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-toolbar-title>Policies</v-toolbar-title>
          <v-divider
              class="mx-4"
              inset
              vertical
          ></v-divider>
          <v-text-field
              v-model="search"
              append-icon="mdi-magnify"
              label="Search"
              single-line
              hide-details
          ></v-text-field>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                  color="primary"
                  dark
                  class="mb-2"
                  v-bind="attrs"
                  v-on="on"
              >New Policy
              </v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-text-field :disabled="editeNameDisabled" v-model="editedItem.metadata.name"
                                label="Name"></v-text-field>
                  <v-text-field v-model="editedItem.policy.description" label="Description"></v-text-field>
                  <v-combobox
                      v-model="editedItem.policy.subjects"
                      hide-selected
                      label="Subjects"
                      multiple
                      small-chips
                  >
                  </v-combobox>
                  <v-select
                      v-model="editedItem.policy.effect"
                      :items="policyEffects"
                      label="Effect"
                  ></v-select>

                  <v-combobox
                      v-model="editedItem.policy.resources"
                      hide-selected
                      label="Resources"
                      multiple
                      small-chips
                  >
                  </v-combobox>

                  <v-combobox
                      v-model="editedItem.policy.actions"
                      hide-selected
                      :items="actionCandidates"
                      :search-input.sync="actionSearch"
                      label="Actions"
                      multiple
                      small-chips
                  >
                  </v-combobox>
                </v-container>

              </v-card-text>

              <v-card-actions>
                <v-checkbox
                    v-model="confirmed"
                    label="Confirm"
                ></v-checkbox>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
                <v-btn :disabled="!confirmed" color="blue darken-1" text @click="save">Save</v-btn>
              </v-card-actions>

            </v-card>
          </v-dialog>
        </v-toolbar>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-icon
            small
            class="mr-2"
            @click="editItem(item)"
        >
          mdi-pencil
        </v-icon>
        <v-icon
            small
            @click="deleteItem(item)"
        >
          mdi-delete
        </v-icon>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="initialize">Refresh</v-btn>
      </template>
    </v-data-table>
    <!--    <v-container class="pa-4">-->
    <!--      <div class="text-center">-->
    <!--        <v-btn fab dark small color="black" @click="initialize()">-->
    <!--          <v-icon>mdi-refresh</v-icon>-->
    <!--        </v-btn>-->

    <!--      </div>-->
    <!--    </v-container>-->


  </div>
</template>

<script>
import axios from "axios";

function getRootPath() {
  return window.location.protocol + '//' + window.location.host;
}

async function createPolicy(policy) {


  let targetURL = getRootPath() + "/v1/admin/policies"
  let succeed = false;
  try {
    await axios.post(targetURL, JSON.stringify(policy)).then(
        response => {
          succeed = response.status === 200;
        }
    )
  } catch (err) {
    succeed = false
  }

  return succeed
}

async function updatePolicy(policy) {
  if (policy === null) {
    return false
  }
  let targetURL = getRootPath() + "/v1/admin/policies/" + policy.metadata.name;
  let succeed = false;
  try {
    await axios.put(targetURL, JSON.stringify(policy)
    ).then(response => {
      if (response.status === 200) {
        succeed = true
      }
    })
  } catch (err) {
    succeed = false
  }

  return succeed
}

async function deletePolicy(policyName) {
  if (policyName === "") {
    return false
  }
  let targetURL = getRootPath() + "/v1/admin/policies/" + policyName;
  let succeed = false;
  try {
    await axios.delete(targetURL
    ).then(response => {
      if (response.status === 200) {
        succeed = true
      }
    })
  } catch (err) {
    succeed = false
  }

  return succeed
}


export default {
  name: "Policies",
  data: () => ({
    search: "",
    confirmed: false,
    policyEffects: ["allow", "deny"],
    headers: [
      {text: 'name', value: 'metadata.name'},
      {text: 'creator', value: 'username'},
      {text: 'description', value: 'policy.description',},
      {text: 'subjects', value: 'policy.subjects'},
      {text: 'effect', value: 'policy.effect'},
      {text: 'resources', value: 'policy.resources'},
      {text: 'actions', value: 'policy.actions'},
      {text: 'Actions', value: 'actions', sortable: false},
    ],
    actionSearch: null,
    actionCandidates: ['get', 'create', 'delete', 'list', 'update', 'watch',"*"],
    policyItems: [],
    editedIndex: -1,
    editedItem: {
      metadata: {
        name: '',
      },
      policy: {
        description: '',
        subjects: [],
        effect: '',
        resources: [],
        actions: [],
      },
    },
    defaultItem: {
      metadata: {
        name: '',
      },
      policy: {
        description: '',
        subjects: [],
        effect: '',
        resources: [],
        actions: [],
      },
    },
    dialog: false,
  }),
  mounted: function () {
    this.setGlobalTitle();
  },
  methods: {
    initialize() {
      let targetURL = getRootPath() + "/v1/admin/policies";
      try {
        axios.get(targetURL,
        ).then(response => {
          // console.log(JSON.stringify(response.data.items, null, 4))
          if (response.status === 200) {
            this.policyItems = response.data.items
          } else {
            this.policyItems = []
          }
        })
      } catch (err) {
        this.policyItems = []
      }
    },
    editItem(item) {
      this.editedIndex = this.policyItems.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
      // console.log(this.editedIndex, this.editedItem)
    },

    async deleteItem(item) {
      const index = this.policyItems.indexOf(item)
      let ret = confirm('Are you sure you want to delete this item?')
      if (ret === true) {
        let succeed = await deletePolicy(this.policyItems[index].metadata.name)
        if (succeed) {
          this.$message.bottom().success("Policy Deleted")
          this.policyItems.splice(index, 1)
        } else {
          this.$message.bottom().success("Failed to Delete Policy")

        }
      }
    },

    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
        // console.log(this.editedItem)
      })
    },

    async save() {
      let succeed = false
      if (this.editedIndex > -1) {
        Object.assign(this.policyItems[this.editedIndex], this.editedItem)
        succeed = await updatePolicy(this.editedItem)
        if (succeed) {
          this.$message.bottom().success("Policy Updated")
        } else {
          this.$message.bottom().success("Failed to Update Policy")
        }

      } else {
        this.policyItems.push(this.editedItem)
        succeed = await createPolicy(this.editedItem)
        if (succeed) {
          this.$message.bottom().success("Policy Created")
        } else {
          this.$message.bottom().success("Failed to Create Policy")
        }
      }

      this.close()
    },
    setGlobalTitle: function () {
      let myElement = document.getElementById("global-title");
      myElement.textContent = "Policies";
    },
  },
  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Policy' : 'Edit Policy'
    },
    editeNameDisabled() {
      return this.editedIndex !== -1;
    }
  },
  watch: {
    dialog(val) {
      val || this.close()
    },
  },
  created() {
    this.initialize()
  },

};
</script>
