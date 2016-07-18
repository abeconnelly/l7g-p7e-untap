var query = [

  "select s0.human_id,",
  "  s0.phenotype_category,",
  "  s0.phenotype,",
  "  s1.phenotype_category,",
  "  s1.phenotype",
  "from survey s0, survey s1",
  "where s0.phenotype_category = 'Basic_Phenotypes:Blood Type' and s0.phenotype = 'AB +'",
  "  and s1.phenotype_category = 'Nervous_System' and s1.phenotype like '%sleep paralysis%'",
  "  and s0.human_id = s1.human_id",

""].join("\n");

var r = pheno_sql(query);
var r_json = JSON.parse(r);

pheno_return(r_json, " ");
