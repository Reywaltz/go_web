//Получение предметов
function getSubjects(data) {
    let tmp_res = [];
    for (let index = 0; index < data.length; index++) {
      tmp_res.push(data[index].short_name);
    }
    let result = [...new Set(tmp_res)];
    return result;
}

//Итоговый массив
function createFinalData(data, subjects) {
    let result = [];
    for (let index = 0; index < data.length; index++) {
      let new_elem = {
        id: data[index].id,
        ФИО: `${data[index].surname} ${data[index].name} ${data[index].second_name}`,
        Группа: data[index].study_group
      };
      for (let subj of subjects) {
        if (data[index].short_name === `${subj}`) {
          new_elem[`${subj}`] = data[index].mark_full;
        } else {
          new_elem[`${subj}`] = "-";
        }
      }
      result.push(new_elem);
    }
    result = filterPairds(
      createPairs_arr(result),
      subjects
    );
    return result;
}


//Разбиение массива на подмассивы по id
function createPairs_arr(arr) {
    let tmp = new Set();
    for (let index = 0; index < arr.length; index++) {
      tmp.add(arr[index]["id"]);
    }
    let allId = [...tmp];
    let result = {};
    for (let el of allId) {
      result[`${el}`] = [];
    }
    for (let index = 0; index < arr.length; index++) {
      for (let id of allId) {
        if (arr[index]["id"] === id) {
          result[`${id}`].push(arr[index]);
        }
      }
    }
    return result;
  }
  //Обновление записей в подмассивах в соответствии с записями по предметам
function filterPairds(arr) {
    let result = [];
    for (let key in arr) {
      if (arr[`${key}`].length === 1) {
        result.push(...arr[`${key}`]);
      } else {
        result.push(middlefilter(arr[`${key}`]));
      }
    }
    return result;
  }

  // initArray = начальный массив данных
function middlefilter(initArray) {
    let newArrData = initArray[0];
    for (let index = 0; index < initArray.length; index++) {
      Object.keys(newArrData).forEach((el, i) => {
        if (
          Object.values(newArrData)[i] !== Object.values(initArray[index])[i] &&
          Object.values(newArrData)[i] === "-"
        ) {
          newArrData = {
            ...newArrData,
            [Object.keys(newArrData)[i]]: Object.values(initArray[index])[i],
          };
        }
      });
    }
    return newArrData;
}


const table = document.querySelector('.styled-table');
const button_all_students = document.querySelector('.get_students');
const button_all_subjects = document.querySelector('.get_subjects');
const button_all_debts = document.querySelector('.get_debts')

let requestOptions = {
    method: 'GET'
};



async function /*Promise<json>*/ getJSONAsync(url) {
    let response = await fetch(url);
    if (response.ok) {
        let json = await response.json();
        return json;
    }
    else throw new Error(`${response.status}: ${response.statusText}`);
}

//Кнопка вывода всех студентов
button_all_students.addEventListener('click', (e) => {
    table.innerHTML = '';
    let arr = getJSONAsync('http://localhost:8080/student/mark');
    arr.then(res => {
        let tmp = [...res];
        let filter_lst = createFinalData(res, getSubjects(tmp));
        let table_head = Object.keys(filter_lst[0]);
        table.insertAdjacentHTML('afterBegin',bodyFill(filter_lst));
        let table_td = table.getElementsByTagName('tr');
        for(let index = 1; index < table_td.length; ++index){
            for(let td of table_td[index].cells){
                if((td.innerText === 'Неудовлетворительно') ||
                   (td.innerText === 'Незачет') ||
                   (td.innerText === 'Неявка')){
                    td.style.background='red';
                    td.style.color='white';
                };
            }
        }
    });
})

//Кнопка аттестаций
button_all_subjects.addEventListener('click', (e) => {
    table.innerHTML = '';
    let arr = getJSONAsync('http://localhost:8080/subject');
    arr.then(res => {
        table.insertAdjacentHTML('afterBegin', bodyFill(res));
    });
})

//Кнопка пересдач
button_all_debts.addEventListener('click', (e) => {
    table.innerHTML = '';
    let arr = getJSONAsync('http://localhost:8080/student/debts');
    arr.then(res => {
        table.insertAdjacentHTML('afterBegin', bodyFill(res));
    });
})

//Заполнение тела
function bodyFill(arr){
    let result = '';
    let headers = new Set();
    for(let index = 0; index < arr.length; index++){
        for(let key in arr[index]){
            headers.add(key);
        }
    }
    for(let th of [...headers]){
        result += `<th>${th}</th>`
    }
    for(let index = 0; index < arr.length; index++){
        result += '<tr>'
        for(let key in arr[index]){
            result += `<td>${arr[index][key]}</td>`
        }
        result += '</tr>'
    }
    return result;
}