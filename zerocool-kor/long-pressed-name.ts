import Debug from "debug";
const debug = Debug("long-pressed-name");

const isLongPressedName = (name:string, typed:string):boolean => {
  debug(`isLongPressedName started (name:${name}, typed:${typed})`);
  let isLongPressedNameIndex = 0;
  if(!name || !typed){
    return false;
  }
  const a = name.split('');
  const b = typed.split('');
  let pre;
  return !(b.find((c,i) => {
    debug(`[i:${i}] -> ${c}`);    
    if(a[0]===c){
      pre = a.shift();
    }else if(pre===c){
      /* OK */
    }else{
      return true;
    }
  }));
}

debug(`result -> ${isLongPressedName('aalex','aaleex')}`);
debug(`result -> ${isLongPressedName('saeed','ssaaedd')}`);
debug(`result -> ${isLongPressedName('leelee','llllleeelee')}`);
debug(`result -> ${isLongPressedName('laiden','laiden')}`);
debug(`result -> ${isLongPressedName('abbbbcd','aaaaaaaaaaabbbcd')}`);