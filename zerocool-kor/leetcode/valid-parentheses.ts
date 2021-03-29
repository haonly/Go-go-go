const isValid = (s:string): boolean => {
    const openParentheses = ['(','{','['];
    const pmap = {
        '(': 1,
        '{' : 2,
        '[' : 3,
        ']' : -3,
        '}' : -2,
        ')' : -1
    }
    console.log(`isValid(${s}) started`);
    const openList = [];
    for(const c of s.split('')){
        const n = pmap[c];
        if(n===undefined){
            return false;
        }
        if(openParentheses.indexOf(c)>=0){
            openList.push(n);       
        }else if(openList.pop() + n !== 0){
            return false;
        }       
    }
    return true;
}

console.log('result -> ' + isValid('abcd'));
console.log('result -> ' + isValid('()'));
console.log('result -> ' + isValid('(){}[]({[]})'));

