source ../.env

rm heroku_cnames.txt
#get heroku domains
heroku_cnames=$(heroku domains | grep CNAME | awk '{print $1","$3}' )
echo $heroku_cnames > heroku_cnames.txt

for domain_record in `cat heroku_cnames.txt`; do
     domain_name=$(echo $domain_record | cut -d"," -f1 )
     domain_account=$(echo $domain_name | awk '{print substr($0, 3)}')
     dns_target=$(echo $domain_record | cut -d"," -f2)

     #get zone_id
     ZONE_ID=$(curl "https://api.cloudflare.com/client/v4/zones?name=$domain_account" \
          -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
          | jq '.result[0].id' | tr -d '"')


     #get dns records for deletion
     dns_json=$(curl "https://api.cloudflare.com/client/v4/zones/$ZONE_ID/dns_records" \
        -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" | jq '.' )

     results_count=$(<<< $dns_json | jq '.result_info.count')

     #delete dns records in loop
     for ((i=0; i < $results_count ; i++)); do
          dns_record_type=$(jq ".result[$i].type" <<< $dns_json | tr -d '"')
          dns_record_content=$(jq ".result[$i].content" <<< $dns_json | tr -d '"')
          DNS_RECORD_ID=$(jq ".result[$i].id" <<< $dns_json | tr -d '"')

          if [ "$dns_record_type" = "CNAME" ]; then
               heorukudns=$(cut -d"." -f2 <<< $dns_record_content)
               if [ "$heorukudns" = "herokudns" ]; then
                    curl https://api.cloudflare.com/client/v4/zones/$ZONE_ID/dns_records/$DNS_RECORD_ID \
                         -X DELETE \
                         -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN"
               fi
          fi
     done

     #add dns record
     curl "https://api.cloudflare.com/client/v4/zones/$ZONE_ID/dns_records" \
          -X POST \
          -H 'Content-Type: application/json' \
          -H "Authorization: Bearer $CLOUDFLARE_API_TOKEN" \
          -d '{
               "comment": "Domain record added autonatically by heroku deployment",
               "content": "'$dns_target'",
               "name": "'$domain_account'",
               "proxied": true,
               "ttl": 1,
               "type": "CNAME"
          }'

done

rm heroku_cnames.txt
