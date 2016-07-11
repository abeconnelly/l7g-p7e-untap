print = ((typeof(print)==="undefined") ? console.log : print);

// Take the input and Stringify is if it's a JSON object.
// If it's a string or number, return it.
// Otherwise return an empty string.
//
function pheno_return(q, indent) {
  indent = ((typeof(indent)==="undefined") ? '' : indent);
  if (typeof(q)==="undefined") { return ""; }
  if (typeof(q)==="object") {
    var s = "";
    try {
      s = JSON.stringify(q, null, indent);
    } catch(err) {
    }
    return s;
  }
  if (typeof(q)==="string") { return q; }
  if (typeof(q)==="number") { return q; }
  return "";
}

function help() {
  print("Lightning Phenotype Server (untap)");
  return "Lightning Phenotype Server (untap)";
}

function info() {
  return pheno_return({"status":"ok"});
}
