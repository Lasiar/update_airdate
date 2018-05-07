new Vue({
    el: "#one",
    data: {
        date_one_start: "",
        date_one_finish: "",
        answer: "",
        loading: false
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
    methods: {
        sendPB: function () {
            this.loading = true;
            axios.post('/update', {
                who: "pb",
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

        },
    }
});
new Vue({
    el: "#two",
    data: {
        date_one_start: "",
        date_one_finish: "",
        answer: "",
        loading: false
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
    methods: {
        sendAll: function () {
            this.loading = true;
            axios.post('/update', {
                who: "all",
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

        },
    }
})