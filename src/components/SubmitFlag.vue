<template>
    <v-card class="mx-auto" v-if="info !== null">
        <v-snackbar v-model="snackBarVisible" :color="snackBarColor" :timeout="3000" :top="true">{{ message }}
        </v-snackbar>

        <v-card-title>
            <span style="margin-right: 20px;">{{$t('flag.submit_flag')}}</span>
            <v-text-field
                    @keyup.enter.native="submitFlag"
                    v-model="inputFlag"
                    :label="$t('flag.input_your_flag')"
                    clearable
            ></v-text-field>
            <v-btn color="success" style="margin-left: 20px;" @click="submitFlag">
                {{ $t('flag.submit') }}
            </v-btn>
        </v-card-title>
    </v-card>
</template>

<script>
    import axios from "axios";

    export default {
        name: "SubmitFlag",

        data: () => ({
            info: null,
            inputFlag: '',

            message: '',
            snackBarVisible: false,
            snackBarColor: 'success',
        }),

        mounted() {
            this.getInfo()
        },

        methods: {
            getInfo() {
                this.utils.GET("/team/info").then(res => {
                    this.info = res
                })
            },
            submitFlag() {
                axios.post(this.utils.baseURL + '/flag', {
                    'flag': this.inputFlag
                }, {
                    headers: {
                        'Authorization': this.info.Token
                    }
                }).then(res => {
                    this.inputFlag = ''
                    this.message = res.data.data
                    this.snackBarColor = 'success'
                    this.snackBarVisible = true
                }).catch(err => {
                    this.message = err.response.data.msg
                    this.snackBarColor = 'error'
                    this.snackBarVisible = true
                });
            }
        }
    }
</script>

<style scoped>

</style>