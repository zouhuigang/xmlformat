package main

import (
	"fmt"
	//"github.com/zouhuigang/xmlformat"
	"./xmlformat"
)

func main2() {
	xmlstring := `<?xml version="1.0" encoding="UTF-8"?><Book><Page><Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1"><Packet x="171.816298" y="143.266575" pressure="0.000000" deltaTime="0.000000"></Packet><Packet x="183.791436" y="191.167127" pressure="120.000000" deltaTime="80.000000"></Packet><Packet x="194.357735" y="225.683702" pressure="122.000000" deltaTime="189.333334"></Packet><Packet x="194.357735" y="229.205801" pressure="122.000000" deltaTime="333.688889"></Packet><Packet x="138.004144" y="213.004144" pressure="118.000000" deltaTime="578.888889"></Packet><Packet x="115.462707" y="212.299724" pressure="120.000000" deltaTime="904.133334"></Packet><Packet x="123.211326" y="213.708564" pressure="122.000000" deltaTime="1377.315556"></Packet><Packet x="135.890884" y="210.890884" pressure="122.000000" deltaTime="1855.582222"></Packet><Packet x="239.440608" y="179.191989" pressure="122.000000" deltaTime="2385.093333"></Packet><Packet x="321.153315" y="167.921271" pressure="122.000000" deltaTime="3000.520000"></Packet><Packet x="294.385359" y="165.808011" pressure="120.000000" deltaTime="3696.786667"></Packet><Packet x="285.932320" y="163.694751" pressure="120.000000" deltaTime="4395.066667"></Packet><Packet x="279.592541" y="159.468232" pressure="112.000000" deltaTime="5104.893334"></Packet><Packet x="270.435083" y="143.266575" pressure="64.000000" deltaTime="5840.048889"></Packet><Packet x="269.026243" y="131.995856" pressure="0.000000" deltaTime="6584.662223"></Packet><Packet x="268.321823" y="122.838398" pressure="30.000000" deltaTime="7340.662222"></Packet><Packet x="270.435083" y="112.976519" pressure="100.000000" deltaTime="8121.986667"></Packet><Packet x="277.479282" y="108.750000" pressure="118.000000" deltaTime="8952.768890"></Packet><Packet x="278.183702" y="110.158840" pressure="118.000000" deltaTime="9816.604445"></Packet><Packet x="273.252762" y="147.493094" pressure="120.000000" deltaTime="10782.342223"></Packet><Packet x="265.504144" y="172.852210" pressure="122.000000" deltaTime="11760.991111"></Packet><Packet x="251.415746" y="203.846685" pressure="122.000000" deltaTime="12762.515556"></Packet><Packet x="236.622928" y="228.501381" pressure="122.000000" deltaTime="13787.693333"></Packet><Packet x="225.352210" y="240.476519" pressure="122.000000" deltaTime="14836.386667"></Packet><Packet x="220.421271" y="243.998619" pressure="122.000000" deltaTime="15894.600000"></Packet><Packet x="211.968232" y="247.520718" pressure="118.000000" deltaTime="16964.200000"></Packet><Packet x="202.106354" y="251.747238" pressure="66.000000" deltaTime="18046.124444"></Packet></Stroke><Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1"><Packet x="176.042818" y="263.722376" pressure="0.000000" deltaTime="0.000000"></Packet><Packet x="176.042818" y="263.017956" pressure="108.000000" deltaTime="13.000000"></Packet><Packet x="178.860497" y="263.017956" pressure="120.000000" deltaTime="64.266666"></Packet><Packet x="188.722376" y="270.766575" pressure="122.000000" deltaTime="150.257778"></Packet><Packet x="195.766575" y="277.810773" pressure="122.000000" deltaTime="271.093333"></Packet><Packet x="195.062155" y="280.628453" pressure="122.000000" deltaTime="413.773333"></Packet><Packet x="176.747238" y="293.308011" pressure="110.000000" deltaTime="593.031110"></Packet><Packet x="146.457182" y="310.214088" pressure="0.000000" deltaTime="793.075555"></Packet></Stroke><Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1"><Packet x="121.098066" y="329.233425" pressure="0.000000" deltaTime="0.000000"></Packet><Packet x="114.053867" y="334.868785" pressure="100.000000" deltaTime="13.000000"></Packet><Packet x="107.714088" y="341.208564" pressure="118.000000" deltaTime="51.266666"></Packet><Packet x="106.305249" y="344.026243" pressure="122.000000" deltaTime="125.991110"></Packet><Packet x="106.305249" y="346.139503" pressure="122.000000" deltaTime="288.502222"></Packet><Packet x="157.023481" y="322.893646" pressure="122.000000" deltaTime="531.791111"></Packet><Packet x="166.180939" y="323.598066" pressure="122.000000" deltaTime="817.093334"></Packet><Packet x="166.180939" y="409.537293" pressure="122.000000" deltaTime="1188.608890"></Packet><Packet x="160.545580" y="474.343923" pressure="124.000000" deltaTime="1627.964445"></Packet><Packet x="202.810773" y="365.158840" pressure="122.000000" deltaTime="2190.066667"></Packet><Packet x="217.603591" y="322.189227" pressure="118.000000" deltaTime="2748.791112"></Packet><Packet x="233.100829" y="277.106354" pressure="8.000000" deltaTime="3319.240001"></Packet></Stroke><Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1"><Packet x="235.214088" y="255.269337" pressure="0.000000" deltaTime="0.000000"></Packet><Packet x="234.509669" y="255.973757" pressure="104.000000" deltaTime="13.000000"></Packet><Packet x="235.214088" y="260.904696" pressure="120.000000" deltaTime="64.266666"></Packet><Packet x="236.622928" y="267.244475" pressure="120.000000" deltaTime="123.257778"></Packet><Packet x="237.327348" y="279.219613" pressure="120.000000" deltaTime="194.693333"></Packet><Packet x="238.031768" y="291.899171" pressure="120.000000" deltaTime="277.320000"></Packet><Packet x="238.736188" y="303.874309" pressure="122.000000" deltaTime="371.275556"></Packet><Packet x="239.440608" y="314.440608" pressure="122.000000" deltaTime="477.555556"></Packet><Packet x="240.145028" y="320.780387" pressure="122.000000" deltaTime="595.026667"></Packet><Packet x="240.145028" y="322.189227" pressure="122.000000" deltaTime="723.826666"></Packet><Packet x="240.145028" y="322.189227" pressure="122.000000" deltaTime="864.951111"></Packet><Packet x="240.849448" y="285.559392" pressure="118.000000" deltaTime="1084.266666"></Packet><Packet x="242.962707" y="274.993094" pressure="116.000000" deltaTime="1305.977777"></Packet><Packet x="252.824586" y="255.269337" pressure="116.000000" deltaTime="1566.311110"></Packet><Packet x="267.617403" y="243.294199" pressure="118.000000" deltaTime="1861.368889"></Packet><Packet x="303.542818" y="226.388122" pressure="122.000000" deltaTime="2311.271111"></Packet><Packet x="285.932320" y="322.893646" pressure="124.000000" deltaTime="2940.017777"></Packet><Packet x="281.705801" y="329.937845" pressure="124.000000" deltaTime="3582.808888"></Packet><Packet x="272.548343" y="328.529006" pressure="124.000000" deltaTime="4261.155555"></Packet><Packet x="252.824586" y="317.962707" pressure="120.000000" deltaTime="4961.346666"></Packet><Packet x="240.145028" y="296.125691" pressure="114.000000" deltaTime="5684.115555"></Packet><Packet x="236.622928" y="284.150552" pressure="114.000000" deltaTime="6430.537777"></Packet><Packet x="241.553867" y="279.219613" pressure="118.000000" deltaTime="7213.475555"></Packet><Packet x="270.435083" y="272.879834" pressure="120.000000" deltaTime="8044.199999"></Packet><Packet x="283.819061" y="271.470994" pressure="122.000000" deltaTime="8895.035555"></Packet><Packet x="292.976519" y="274.288674" pressure="122.000000" deltaTime="9809.506665"></Packet><Packet x="282.410221" y="289.785912" pressure="120.000000" deltaTime="10768.164444"></Packet><Packet x="240.145028" y="334.868785" pressure="122.000000" deltaTime="11827.053332"></Packet><Packet x="300.020718" y="310.214088" pressure="122.000000" deltaTime="13044.911110"></Packet><Packet x="290.863260" y="315.145028" pressure="122.000000" deltaTime="14320.177777"></Packet><Packet x="264.095304" y="343.321823" pressure="122.000000" deltaTime="15638.546665"></Packet><Packet x="201.401934" y="425.034530" pressure="122.000000" deltaTime="17030.204443"></Packet><Packet x="198.584254" y="425.738950" pressure="124.000000" deltaTime="18505.431110"></Packet><Packet x="199.288674" y="422.921271" pressure="124.000000" deltaTime="19994.479999"></Packet><Packet x="200.697514" y="419.399171" pressure="122.000000" deltaTime="21494.479999"></Packet><Packet x="220.421271" y="397.562155" pressure="122.000000" deltaTime="23072.728888"></Packet><Packet x="247.893646" y="380.656077" pressure="122.000000" deltaTime="24693.373332"></Packet><Packet x="305.656077" y="356.001381" pressure="122.000000" deltaTime="26360.306666"></Packet><Packet x="343.694751" y="349.661602" pressure="124.000000" deltaTime="28074.408888"></Packet><Packet x="342.285912" y="376.429558" pressure="124.000000" deltaTime="30021.546665"></Packet><Packet x="328.901934" y="456.733425" pressure="124.000000" deltaTime="32002.924443"></Packet><Packet x="326.084254" y="468.708564" pressure="124.000000" deltaTime="33990.568887"></Packet><Packet x="314.813536" y="491.250000" pressure="124.000000" deltaTime="36042.639998"></Packet><Packet x="305.656077" y="487.727901" pressure="124.000000" deltaTime="38125.973331"></Packet><Packet x="293.680939" y="476.457182" pressure="124.000000" deltaTime="40231.266664"></Packet><Packet x="266.912983" y="434.896409" pressure="122.000000" deltaTime="42373.137775"></Packet><Packet x="250.006906" y="408.832873" pressure="98.000000" deltaTime="44536.795553"></Packet><Packet x="245.075967" y="401.084254" pressure="0.000000" deltaTime="46710.031109"></Packet></Stroke><Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1"><Packet x="273.957182" y="356.001381" pressure="0.000000" deltaTime="0.000000"></Packet><Packet x="275.366022" y="355.296961" pressure="78.000000" deltaTime="13.000000"></Packet><Packet x="277.479282" y="355.296961" pressure="102.000000" deltaTime="37.266666"></Packet><Packet x="278.183702" y="356.705801" pressure="110.000000" deltaTime="73.857778"></Packet><Packet x="279.592541" y="360.227901" pressure="114.000000" deltaTime="121.640000"></Packet><Packet x="279.592541" y="375.020718" pressure="116.000000" deltaTime="194.751111"></Packet><Packet x="278.183702" y="386.291436" pressure="118.000000" deltaTime="277.320000"></Packet><Packet x="275.366022" y="398.266575" pressure="120.000000" deltaTime="371.275556"></Packet><Packet x="265.504144" y="427.852210" pressure="122.000000" deltaTime="503.555556"></Packet><Packet x="261.277624" y="434.191989" pressure="122.000000" deltaTime="644.560000"></Packet><Packet x="253.529006" y="438.418508" pressure="122.000000" deltaTime="823.875556"></Packet><Packet x="252.824586" y="436.305249" pressure="122.000000" deltaTime="1010.920000"></Packet><Packet x="253.529006" y="422.921271" pressure="118.000000" deltaTime="1223.408889"></Packet><Packet x="261.982044" y="410.946133" pressure="108.000000" deltaTime="1458.355556"></Packet><Packet x="277.479282" y="405.310773" pressure="110.000000" deltaTime="1716.955556"></Packet><Packet x="310.587017" y="405.310773" pressure="120.000000" deltaTime="2052.071112"></Packet><Packet x="292.272099" y="413.059392" pressure="118.000000" deltaTime="2429.640001"></Packet><Packet x="281.705801" y="415.172652" pressure="104.000000" deltaTime="2814.497778"></Packet><Packet x="271.843923" y="415.877072" pressure="0.000000" deltaTime="3210.724445"></Packet></Stroke><Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1"><Packet x="234.509669" y="406.015193" pressure="0.000000" deltaTime="0.000000"></Packet><Packet x="233.100829" y="407.424033" pressure="96.000000" deltaTime="13.000000"></Packet><Packet x="232.396409" y="409.537293" pressure="114.000000" deltaTime="51.266666"></Packet><Packet x="234.509669" y="415.877072" pressure="118.000000" deltaTime="111.991110"></Packet><Packet x="235.918508" y="421.512431" pressure="118.000000" deltaTime="183.368888"></Packet><Packet x="236.622928" y="428.556630" pressure="118.000000" deltaTime="265.995555"></Packet><Packet x="236.622928" y="437.714088" pressure="120.000000" deltaTime="359.951111"></Packet><Packet x="235.918508" y="446.871547" pressure="120.000000" deltaTime="466.231110"></Packet><Packet x="235.214088" y="453.915746" pressure="120.000000" deltaTime="583.702222"></Packet><Packet x="234.509669" y="460.959945" pressure="120.000000" deltaTime="712.502221"></Packet><Packet x="233.805249" y="465.890884" pressure="122.000000" deltaTime="853.626666"></Packet><Packet x="233.805249" y="468.708564" pressure="122.000000" deltaTime="1005.942221"></Packet><Packet x="297.907459" y="446.871547" pressure="122.000000" deltaTime="1263.586666"></Packet><Packet x="316.222376" y="442.645028" pressure="122.000000" deltaTime="1533.022221"></Packet><Packet x="319.744475" y="444.053867" pressure="122.000000" deltaTime="1839.466666"></Packet><Packet x="314.813536" y="444.758287" pressure="124.000000" deltaTime="2154.693332"></Packet><Packet x="307.064917" y="446.871547" pressure="122.000000" deltaTime="2481.231110"></Packet><Packet x="290.863260" y="444.758287" pressure="104.000000" deltaTime="2833.097776"></Packet></Stroke></Page></Book>`
	x1 := xmlformat.FormatXML(xmlstring, " ", "  ")
	fmt.Println(x1)
}

func main() {
	x3 := `<?xml version="1.0" encoding="utf-8"?>


<Book>


  <Page>
    <Stroke duration="1240"   startSecond="1439307219" startMillisecond="517" lineWidth="1">   
      <Packet x="171.816298" y="143.266575" pressure="0.000000" deltaTime="0.000000"/>
      <Packet x="183.791436" y="191.167127" pressure="120.000000" deltaTime="80.000000"/>
      <Packet x="194.357735" y="225.683702" pressure="122.000000" deltaTime="189.333334"/>
      <Packet x="194.357735" y="229.205801" pressure="122.000000" deltaTime="333.688889"/>
      <Packet x="138.004144" y="213.004144" pressure="118.000000" deltaTime="578.888889"/>
      <Packet x="115.462707" y="212.299724" pressure="120.000000" deltaTime="904.133334"/>
      <Packet x="123.211326" y="213.708564" pressure="122.000000" deltaTime="1377.315556"/>
      <Packet x="135.890884" y="210.890884" pressure="122.000000" deltaTime="1855.582222"/>
      <Packet x="239.440608" y="179.191989" pressure="122.000000" deltaTime="2385.093333"/>
      <Packet x="321.153315" y="167.921271" pressure="122.000000" deltaTime="3000.520000"/>
      <Packet x="294.385359" y="165.808011" pressure="120.000000" deltaTime="3696.786667"/>
      <Packet x="285.932320" y="163.694751" pressure="120.000000" deltaTime="4395.066667"/>
      <Packet x="279.592541" y="159.468232" pressure="112.000000" deltaTime="5104.893334"/>
      <Packet x="270.435083" y="143.266575" pressure="64.000000" deltaTime="5840.048889"/>
      <Packet x="269.026243" y="131.995856" pressure="0.000000" deltaTime="6584.662223"/>
      <Packet x="268.321823" y="122.838398" pressure="30.000000" deltaTime="7340.662222"/>
      <Packet x="270.435083" y="112.976519" pressure="100.000000" deltaTime="8121.986667"/>
      <Packet x="277.479282" y="108.750000" pressure="118.000000" deltaTime="8952.768890"/>
      <Packet x="278.183702" y="110.158840" pressure="118.000000" deltaTime="9816.604445"/>
      <Packet x="273.252762" y="147.493094" pressure="120.000000" deltaTime="10782.342223"/>
      <Packet x="265.504144" y="172.852210" pressure="122.000000" deltaTime="11760.991111"/>
      <Packet x="251.415746" y="203.846685" pressure="122.000000" deltaTime="12762.515556"/>
      <Packet x="236.622928" y="228.501381" pressure="122.000000" deltaTime="13787.693333"/>
      <Packet x="225.352210" y="240.476519" pressure="122.000000" deltaTime="14836.386667"/>
      <Packet x="220.421271" y="243.998619" pressure="122.000000" deltaTime="15894.600000"/>
      <Packet x="211.968232" y="247.520718" pressure="118.000000" deltaTime="16964.200000"/>
      <Packet x="202.106354" y="251.747238" pressure="66.000000" deltaTime="18046.124444"/>
    </Stroke>
    <Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1">
      <Packet x="176.042818" y="263.722376" pressure="0.000000" deltaTime="0.000000"/>
      <Packet x="176.042818" y="263.017956" pressure="108.000000" deltaTime="13.000000"/>
      <Packet x="178.860497" y="263.017956" pressure="120.000000" deltaTime="64.266666"/>
      <Packet x="188.722376" y="270.766575" pressure="122.000000" deltaTime="150.257778"/>
      <Packet x="195.766575" y="277.810773" pressure="122.000000" deltaTime="271.093333"/>
      <Packet x="195.062155" y="280.628453" pressure="122.000000" deltaTime="413.773333"/>
      <Packet x="176.747238" y="293.308011" pressure="110.000000" deltaTime="593.031110"/>
      <Packet x="146.457182" y="310.214088" pressure="0.000000" deltaTime="793.075555"/>
    </Stroke>
    <Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1">
      <Packet x="121.098066" y="329.233425" pressure="0.000000" deltaTime="0.000000"/>
      <Packet x="114.053867" y="334.868785" pressure="100.000000" deltaTime="13.000000"/>
      <Packet x="107.714088" y="341.208564" pressure="118.000000" deltaTime="51.266666"/>
      <Packet x="106.305249" y="344.026243" pressure="122.000000" deltaTime="125.991110"/>
      <Packet x="106.305249" y="346.139503" pressure="122.000000" deltaTime="288.502222"/>
      <Packet x="157.023481" y="322.893646" pressure="122.000000" deltaTime="531.791111"/>
      <Packet x="166.180939" y="323.598066" pressure="122.000000" deltaTime="817.093334"/>
      <Packet x="166.180939" y="409.537293" pressure="122.000000" deltaTime="1188.608890"/>
      <Packet x="160.545580" y="474.343923" pressure="124.000000" deltaTime="1627.964445"/>
      <Packet x="202.810773" y="365.158840" pressure="122.000000" deltaTime="2190.066667"/>
      <Packet x="217.603591" y="322.189227" pressure="118.000000" deltaTime="2748.791112"/>
      <Packet x="233.100829" y="277.106354" pressure="8.000000" deltaTime="3319.240001"/>
    </Stroke>
    <Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1">
      <Packet x="235.214088" y="255.269337" pressure="0.000000" deltaTime="0.000000"/>
      <Packet x="234.509669" y="255.973757" pressure="104.000000" deltaTime="13.000000"/>
      <Packet x="235.214088" y="260.904696" pressure="120.000000" deltaTime="64.266666"/>
      <Packet x="236.622928" y="267.244475" pressure="120.000000" deltaTime="123.257778"/>
      <Packet x="237.327348" y="279.219613" pressure="120.000000" deltaTime="194.693333"/>
      <Packet x="238.031768" y="291.899171" pressure="120.000000" deltaTime="277.320000"/>
      <Packet x="238.736188" y="303.874309" pressure="122.000000" deltaTime="371.275556"/>
      <Packet x="239.440608" y="314.440608" pressure="122.000000" deltaTime="477.555556"/>
      <Packet x="240.145028" y="320.780387" pressure="122.000000" deltaTime="595.026667"/>
      <Packet x="240.145028" y="322.189227" pressure="122.000000" deltaTime="723.826666"/>
      <Packet x="240.145028" y="322.189227" pressure="122.000000" deltaTime="864.951111"/>
      <Packet x="240.849448" y="285.559392" pressure="118.000000" deltaTime="1084.266666"/>
      <Packet x="242.962707" y="274.993094" pressure="116.000000" deltaTime="1305.977777"/>
      <Packet x="252.824586" y="255.269337" pressure="116.000000" deltaTime="1566.311110"/>
      <Packet x="267.617403" y="243.294199" pressure="118.000000" deltaTime="1861.368889"/>
      <Packet x="303.542818" y="226.388122" pressure="122.000000" deltaTime="2311.271111"/>
      <Packet x="285.932320" y="322.893646" pressure="124.000000" deltaTime="2940.017777"/>
      <Packet x="281.705801" y="329.937845" pressure="124.000000" deltaTime="3582.808888"/>
      <Packet x="272.548343" y="328.529006" pressure="124.000000" deltaTime="4261.155555"/>
      <Packet x="252.824586" y="317.962707" pressure="120.000000" deltaTime="4961.346666"/>
      <Packet x="240.145028" y="296.125691" pressure="114.000000" deltaTime="5684.115555"/>
      <Packet x="236.622928" y="284.150552" pressure="114.000000" deltaTime="6430.537777"/>
      <Packet x="241.553867" y="279.219613" pressure="118.000000" deltaTime="7213.475555"/>
      <Packet x="270.435083" y="272.879834" pressure="120.000000" deltaTime="8044.199999"/>
      <Packet x="283.819061" y="271.470994" pressure="122.000000" deltaTime="8895.035555"/>
      <Packet x="292.976519" y="274.288674" pressure="122.000000" deltaTime="9809.506665"/>
      <Packet x="282.410221" y="289.785912" pressure="120.000000" deltaTime="10768.164444"/>
      <Packet x="240.145028" y="334.868785" pressure="122.000000" deltaTime="11827.053332"/>
      <Packet x="300.020718" y="310.214088" pressure="122.000000" deltaTime="13044.911110"/>
      <Packet x="290.863260" y="315.145028" pressure="122.000000" deltaTime="14320.177777"/>
      <Packet x="264.095304" y="343.321823" pressure="122.000000" deltaTime="15638.546665"/>
      <Packet x="201.401934" y="425.034530" pressure="122.000000" deltaTime="17030.204443"/>
      <Packet x="198.584254" y="425.738950" pressure="124.000000" deltaTime="18505.431110"/>
      <Packet x="199.288674" y="422.921271" pressure="124.000000" deltaTime="19994.479999"/>
      <Packet x="200.697514" y="419.399171" pressure="122.000000" deltaTime="21494.479999"/>
      <Packet x="220.421271" y="397.562155" pressure="122.000000" deltaTime="23072.728888"/>
      <Packet x="247.893646" y="380.656077" pressure="122.000000" deltaTime="24693.373332"/>
      <Packet x="305.656077" y="356.001381" pressure="122.000000" deltaTime="26360.306666"/>
      <Packet x="343.694751" y="349.661602" pressure="124.000000" deltaTime="28074.408888"/>
      <Packet x="342.285912" y="376.429558" pressure="124.000000" deltaTime="30021.546665"/>
      <Packet x="328.901934" y="456.733425" pressure="124.000000" deltaTime="32002.924443"/>
      <Packet x="326.084254" y="468.708564" pressure="124.000000" deltaTime="33990.568887"/>
      <Packet x="314.813536" y="491.250000" pressure="124.000000" deltaTime="36042.639998"/>
      <Packet x="305.656077" y="487.727901" pressure="124.000000" deltaTime="38125.973331"/>
      <Packet x="293.680939" y="476.457182" pressure="124.000000" deltaTime="40231.266664"/>
      <Packet x="266.912983" y="434.896409" pressure="122.000000" deltaTime="42373.137775"/>
      <Packet x="250.006906" y="408.832873" pressure="98.000000" deltaTime="44536.795553"/>
      <Packet x="245.075967" y="401.084254" pressure="0.000000" deltaTime="46710.031109"/>
    </Stroke>
    <Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1">
      <Packet x="273.957182" y="356.001381" pressure="0.000000" deltaTime="0.000000"/>
      <Packet x="275.366022" y="355.296961" pressure="78.000000" deltaTime="13.000000"/>
      <Packet x="277.479282" y="355.296961" pressure="102.000000" deltaTime="37.266666"/>
      <Packet x="278.183702" y="356.705801" pressure="110.000000" deltaTime="73.857778"/>
      <Packet x="279.592541" y="360.227901" pressure="114.000000" deltaTime="121.640000"/>
      <Packet x="279.592541" y="375.020718" pressure="116.000000" deltaTime="194.751111"/>
      <Packet x="278.183702" y="386.291436" pressure="118.000000" deltaTime="277.320000"/>
      <Packet x="275.366022" y="398.266575" pressure="120.000000" deltaTime="371.275556"/>
      <Packet x="265.504144" y="427.852210" pressure="122.000000" deltaTime="503.555556"/>
      <Packet x="261.277624" y="434.191989" pressure="122.000000" deltaTime="644.560000"/>
      <Packet x="253.529006" y="438.418508" pressure="122.000000" deltaTime="823.875556"/>
      <Packet x="252.824586" y="436.305249" pressure="122.000000" deltaTime="1010.920000"/>
      <Packet x="253.529006" y="422.921271" pressure="118.000000" deltaTime="1223.408889"/>
      <Packet x="261.982044" y="410.946133" pressure="108.000000" deltaTime="1458.355556"/>
      <Packet x="277.479282" y="405.310773" pressure="110.000000" deltaTime="1716.955556"/>
      <Packet x="310.587017" y="405.310773" pressure="120.000000" deltaTime="2052.071112"/>
      <Packet x="292.272099" y="413.059392" pressure="118.000000" deltaTime="2429.640001"/>
      <Packet x="281.705801" y="415.172652" pressure="104.000000" deltaTime="2814.497778"/>
      <Packet x="271.843923" y="415.877072" pressure="0.000000" deltaTime="3210.724445"/>
    </Stroke>
    <Stroke duration="1240" startSecond="1439307219" startMillisecond="517" lineWidth="1">
      <Packet x="234.509669" y="406.015193" pressure="0.000000" deltaTime="0.000000"/>
      <Packet x="233.100829" y="407.424033" pressure="96.000000" deltaTime="13.000000"/>
      <Packet x="232.396409" y="409.537293" pressure="114.000000" deltaTime="51.266666"/>
      <Packet x="234.509669" y="415.877072" pressure="118.000000" deltaTime="111.991110"/>
      <Packet x="235.918508" y="421.512431" pressure="118.000000" deltaTime="183.368888"/>
      <Packet x="236.622928" y="428.556630" pressure="118.000000" deltaTime="265.995555"/>
      <Packet x="236.622928" y="437.714088" pressure="120.000000" deltaTime="359.951111"/>
      <Packet x="235.918508" y="446.871547" pressure="120.000000" deltaTime="466.231110"/>
      <Packet x="235.214088" y="453.915746" pressure="120.000000" deltaTime="583.702222"/>
      <Packet x="234.509669" y="460.959945" pressure="120.000000" deltaTime="712.502221"/>
      <Packet x="233.805249" y="465.890884" pressure="122.000000" deltaTime="853.626666"/>
      <Packet x="233.805249" y="468.708564" pressure="122.000000" deltaTime="1005.942221"/>
      <Packet x="297.907459" y="446.871547" pressure="122.000000" deltaTime="1263.586666"/>
      <Packet x="316.222376" y="442.645028" pressure="122.000000" deltaTime="1533.022221"/>
      <Packet x="319.744475" y="444.053867" pressure="122.000000" deltaTime="1839.466666"/>
      <Packet x="314.813536" y="444.758287" pressure="124.000000" deltaTime="2154.693332"/>
      <Packet x="307.064917" y="446.871547" pressure="122.000000" deltaTime="2481.231110"/>
      <Packet x="290.863260" y="444.758287" pressure="104.000000" deltaTime="2833.097776"/>
    </Stroke>
  </Page>
</Book>`

	x1 := xmlformat.CompressXml(x3)
	fmt.Println(x1)
}
