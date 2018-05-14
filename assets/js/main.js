Vue.component('date-form', {
    template: '#date-form-template',
    data: function () {
        return {
            date_one_start: "",
            date_one_finish: "",
            loading: false
        }
    },
    props: ['who'],
    methods: {
        update: function () {
            this.loading = true;
            axios.post('/update', {
                who: this.who,
                date_start: this.date_one_start,
                date_finish: this.date_one_finish
            })
                .then(response => {
                    if (response.data.success) {
                        alert("Обновленно")
                    } else {
                        alert("Не обновленно")
                    }
                })
                .catch(function (error) {
                    alert('Ошибка! Не могу связаться с API: ' + error)
                })
                .finally(() => this.loading = false)

        }
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
    el: "#app"
});
