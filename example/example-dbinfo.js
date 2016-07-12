var tables = [ "allergies",
    "human_suff_record_map",
    "suff_record",
    "conditions",
    "immunizations",
    "survey",
    "demographics",
    "insuff_record",
    "test_results",
    "enrollment_date",
    "medications",
    "uploaded_data",
    "geographic_information",
    "procedures",
    "human_insuff_record_map",
    "specimens",
    "huid_lightning_dataset_map"
    ];

var result = {};

for (var i=0; i<tables.length; i++) {

  var query = "select * from " + tables[i] + " limit 2;";

  var r = pheno_sql(query);
  var r_json = JSON.parse(r);
  result[tables[i]] = r_json;
}

pheno_return(result, "  ");
