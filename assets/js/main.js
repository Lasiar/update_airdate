window.Event = new Vue();

Vue.component('modal', {
    template: '#modal-template',
    props: ['succes'],
    computed: {
        refresh: function () {
            return this.succes
        }
    }
});

Vue.component('date-form', {
    template: '#date-form-template',
    data: function () {
        return {
            date_one_start: "",
            date_one_finish: "",
            response: "",
            loading: false
        }
    },
    props: ['who'],
    methods: {
        update: function () {
            let self = this;
            this.loading = true;
            axios.post('/update', {
                who: this.who,
                date_start: this.date_one_start,
                date_finish: this.date_one_finish
            })
                .then(function (response) {
                    self.$emit('response', response.data.success);
                })
                .catch(function (error) {
                    alert('Ошибка! Не могу связаться с API: ' + error)
                })
                .finally(() => this.loading = false)

        },
    },

    computed: {
        checkData: function () {
            if (this.date_one_start == '') {
                return false
            }
            if (this.date_one_finish == '') {
                //console.log("in if");
                return true
            }
            var finish = this.date_one_finish.replace('/-/g', '');
            var start = this.date_one_start.replace('/-/g', '');
            return start <= finish

        }
    },
});


new Vue({
    el: "#app",
    data: {
        showModal: false,
        info: "",
        succes: false
    },
    methods: {
        modal: function (succes) {
            if (succes) {
                this.success = "true";
                this.info = "Обновлено"
            } else {
                this.success = "false";
                this.info = "не обновлено"
            }
            this.showModal = true
        }
    }
});
