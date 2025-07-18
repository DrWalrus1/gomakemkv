package gomakemkv_test

import (
	"strings"
	"testing"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/go-playground/assert/v2"
)

func TestInfoEventHandler(t *testing.T) {
	input :=
		`MSG:1004,0,1,"MakeMKV v1.18.1 linux(x64-release) started","%1 started","MakeMKV v1.18.1 linux(x64-release)"
MSG:3338,0,2,"Downloading latest SDF to /root/.MakeMKV ...","Downloading latest %1 to %2 ...","SDF","/root/.MakeMKV"
DRV:0,2,999,12,"BD-RE ASUS BW-16D1HT 3.10 KL7L9D92707","BDROM","/dev/sr0"
DRV:1,256,999,0,"","",""
DRV:2,256,999,0,"","",""
DRV:3,256,999,0,"","",""
DRV:4,256,999,0,"","",""
DRV:5,256,999,0,"","",""
DRV:6,256,999,0,"","",""
DRV:7,256,999,0,"","",""
DRV:8,256,999,0,"","",""
DRV:9,256,999,0,"","",""
DRV:10,256,999,0,"","",""
DRV:11,256,999,0,"","",""
DRV:12,256,999,0,"","",""
DRV:13,256,999,0,"","",""
DRV:14,256,999,0,"","",""
DRV:15,256,999,0,"","",""
MSG:1011,0,1,"Using LibreDrive mode (v06.3 id=0FA242DD4D0B)","%1","Using LibreDrive mode (v06.3 id=0FA242DD4D0B)"
MSG:3007,0,0,"Using direct disc access mode","Using direct disc access mode"
MSG:5053,1544,1,"This functionality is shareware. You may evaluate it for 30 days after what you would need to purchase an activation key if you like the functionality. Do you want to start evaluation period now?","This functionality is shareware. You may evaluate it for %1 days after what you would need to purchase an activation key if you like the functionality. Do you want to start evaluation period now?","30"
MSG:5050,0,2,"Evaluation version, 30 day(s) out of 30 remaining","Evaluation version, %1 day(s) out of %2 remaining","30","30"
MSG:5085,0,0,"Loaded content hash table, will verify integrity of M2TS files.","Loaded content hash table, will verify integrity of M2TS files."
MSG:3025,0,3,"Title #00001.mpls has length of 12 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00001.mpls","12","120"
MSG:3307,16777216,2,"File 00000.mpls was added as title #0","File %1 was added as title #%2","00000.mpls","0"
MSG:3025,16777216,3,"Title #00017.m2ts has length of 1 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00017.m2ts","1","120"
MSG:3025,16777216,3,"Title #00011.m2ts has length of 1 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00011.m2ts","1","120"
MSG:3025,16777216,3,"Title #00005.m2ts has length of 7 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00005.m2ts","7","120"
MSG:3025,16777216,3,"Title #00004.m2ts has length of 5 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00004.m2ts","5","120"
MSG:3025,16777216,3,"Title #00006.m2ts has length of 9 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00006.m2ts","9","120"
MSG:3025,0,3,"Title #00010.m2ts has length of 39 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00010.m2ts","39","120"
MSG:5011,0,0,"Operation successfully completed","Operation successfully completed"
TCOUNT:1
CINFO:1,6209,"Blu-ray disc"
CINFO:2,0,"Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Standard Edition Disc 1"
CINFO:28,0,"eng"
CINFO:29,0,"English"
CINFO:30,0,"Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Standard Edition Disc 1"
CINFO:31,6119,"<b>Source information</b><br>"
CINFO:32,0,"BDROM"
CINFO:33,0,"0"
TINFO:0,2,0,"Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Standard Edition Disc 1"
TINFO:0,8,0,"21"
TINFO:0,9,0,"1:37:28"
TINFO:0,10,0,"29.6 GB"
TINFO:0,11,0,"31843694592"
TINFO:0,16,0,"00000.mpls"
TINFO:0,25,0,"4"
TINFO:0,26,0,"0,1,2,3"
TINFO:0,27,0,"Demon Slayer- Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Standard Edition Disc 1_t00.mkv"
TINFO:0,28,0,"eng"
TINFO:0,29,0,"English"
TINFO:0,30,0,"Demon Slayer: Kimetsu no Yaiba Swordsmith Village Arc Complete Blu-ray Set Standard Edition Disc 1 - 21 chapter(s) , 29.6 GB"
TINFO:0,31,6120,"<b>Title information</b><br>"
TINFO:0,33,0,"0"
SINFO:0,0,1,6201,"Video"
SINFO:0,0,5,0,"V_MPEG4/ISO/AVC"
SINFO:0,0,6,0,"Mpeg4"
SINFO:0,0,7,0,"Mpeg4 AVC High@L4.1"
SINFO:0,0,19,0,"1920x1080"
SINFO:0,0,20,0,"16:9"
SINFO:0,0,21,0,"23.976 (24000/1001)"
SINFO:0,0,22,0,"0"
SINFO:0,0,28,0,"eng"
SINFO:0,0,29,0,"English"
SINFO:0,0,30,0,"Mpeg4 AVC High@L4.1"
SINFO:0,0,31,6121,"<b>Track information</b><br>"
SINFO:0,0,33,0,"0"
SINFO:0,0,38,0,""
SINFO:0,0,42,5088,"( Lossless conversion )"
SINFO:0,1,1,6202,"Audio"
SINFO:0,1,2,5091,"Stereo"
SINFO:0,1,3,0,"eng"
SINFO:0,1,4,0,"English"
SINFO:0,1,5,0,"A_PCM/INT/LIT"
SINFO:0,1,6,0,"LPCM"
SINFO:0,1,7,0,"Linear PCM"
SINFO:0,1,13,0,"36 Kb/s"
SINFO:0,1,14,0,"2"
SINFO:0,1,17,0,"48000"
SINFO:0,1,18,0,"24"
SINFO:0,1,22,0,"0"
SINFO:0,1,28,0,"eng"
SINFO:0,1,29,0,"English"
SINFO:0,1,30,0,"LPCM Stereo English"
SINFO:0,1,31,6121,"<b>Track information</b><br>"
SINFO:0,1,33,0,"90"
SINFO:0,1,34,0,"Raw LPCM audio without channel configuration metadata ( Lossless conversion )"
SINFO:0,1,35,0,"Save as raw LPCM"
SINFO:0,1,38,0,"d"
SINFO:0,1,39,0,"Default"
SINFO:0,1,40,0,"stereo"
SINFO:0,1,42,5088,"( Lossless conversion )"
SINFO:0,2,1,6202,"Audio"
SINFO:0,2,2,5091,"Stereo"
SINFO:0,2,3,0,"jpn"
SINFO:0,2,4,0,"Japanese"
SINFO:0,2,5,0,"A_PCM/INT/LIT"
SINFO:0,2,6,0,"LPCM"
SINFO:0,2,7,0,"Linear PCM"
SINFO:0,2,13,0,"36 Kb/s"
SINFO:0,2,14,0,"2"
SINFO:0,2,17,0,"48000"
SINFO:0,2,18,0,"24"
SINFO:0,2,22,0,"0"
SINFO:0,2,28,0,"eng"
SINFO:0,2,29,0,"English"
SINFO:0,2,30,0,"LPCM Stereo Japanese"
SINFO:0,2,31,6121,"<b>Track information</b><br>"
SINFO:0,2,33,0,"90"
SINFO:0,2,34,0,"Raw LPCM audio without channel configuration metadata ( Lossless conversion )"
SINFO:0,2,35,0,"Save as raw LPCM"
SINFO:0,2,38,0,""
SINFO:0,2,40,0,"stereo"
SINFO:0,2,42,5088,"( Lossless conversion )"
SINFO:0,3,1,6203,"Subtitles"
SINFO:0,3,3,0,"eng"
SINFO:0,3,4,0,"English"
SINFO:0,3,5,0,"S_HDMV/PGS"
SINFO:0,3,6,0,"PGS"
SINFO:0,3,7,0,"HDMV PGS Subtitles"
SINFO:0,3,22,0,"0"
SINFO:0,3,28,0,"eng"
SINFO:0,3,29,0,"English"
SINFO:0,3,30,0,"PGS English"
SINFO:0,3,31,6121,"<b>Track information</b><br>"
SINFO:0,3,33,0,"90"
SINFO:0,3,38,0,""
SINFO:0,3,42,5088,"( Lossless conversion )"
SINFO:0,4,1,6203,"Subtitles"
SINFO:0,4,3,0,"eng"
SINFO:0,4,4,0,"English"
SINFO:0,4,5,0,"S_HDMV/PGS"
SINFO:0,4,6,0,"PGS"
SINFO:0,4,7,0,"HDMV PGS Subtitles"
SINFO:0,4,22,0,"6144"
SINFO:0,4,28,0,"eng"
SINFO:0,4,29,0,"English"
SINFO:0,4,30,0,"PGS English  (forced only)"
SINFO:0,4,31,6121,"<b>Track information</b><br>"
SINFO:0,4,33,0,"90"
SINFO:0,4,38,0,"d"
SINFO:0,4,39,0,"Default"
SINFO:0,4,42,5088,"( Lossless conversion )"
SINFO:0,5,1,6203,"Subtitles"
SINFO:0,5,3,0,"eng"
SINFO:0,5,4,0,"English"
SINFO:0,5,5,0,"S_HDMV/PGS"
SINFO:0,5,6,0,"PGS"
SINFO:0,5,7,0,"HDMV PGS Subtitles"
SINFO:0,5,22,0,"0"
SINFO:0,5,28,0,"eng"
SINFO:0,5,29,0,"English"
SINFO:0,5,30,0,"PGS English"
SINFO:0,5,31,6121,"<b>Track information</b><br>"
SINFO:0,5,33,0,"90"
SINFO:0,5,38,0,""
SINFO:0,5,42,5088,"( Lossless conversion )"
SINFO:0,6,1,6203,"Subtitles"
SINFO:0,6,3,0,"eng"
SINFO:0,6,4,0,"English"
SINFO:0,6,5,0,"S_HDMV/PGS"
SINFO:0,6,6,0,"PGS"
SINFO:0,6,7,0,"HDMV PGS Subtitles"
SINFO:0,6,22,0,"6144"
SINFO:0,6,28,0,"eng"
SINFO:0,6,29,0,"English"
SINFO:0,6,30,0,"PGS English  (forced only)"
SINFO:0,6,31,6121,"<b>Track information</b><br>"
SINFO:0,6,33,0,"90"
SINFO:0,6,38,0,""
SINFO:0,6,42,5088,"( Lossless conversion )"`

	standardEvents, discInfoEvents := gomakemkv.ParseMakeMkvInfoCommandLogs(strings.NewReader(input))

	standardEventCount := 0
	discInfoEventCount := 0
loop:
	for {
		select {
		case <-standardEvents:
			standardEventCount++
		case <-discInfoEvents:
			discInfoEventCount++
			break loop
		}
	}
	assert.Equal(t, 33, standardEventCount)
	assert.Equal(t, 1, discInfoEventCount)
}

func TestGetInfoWithNoDiscStillExits(t *testing.T) {

	input :=
		`MSG:1004,0,1,"MakeMKV v1.18.1 linux(x64-release) started","%1 started","MakeMKV v1.18.1 linux(x64-release)"
MSG:3338,0,2,"Downloading latest SDF to /root/.MakeMKV ...","Downloading latest %1 to %2 ...","SDF","/root/.MakeMKV"
DRV:0,2,999,12,"BD-RE ASUS BW-16D1HT 3.10 KL7L9D92707","BDROM","/dev/sr0"
DRV:1,256,999,0,"","",""
DRV:2,256,999,0,"","",""
DRV:3,256,999,0,"","",""
DRV:4,256,999,0,"","",""
DRV:5,256,999,0,"","",""
DRV:6,256,999,0,"","",""
DRV:7,256,999,0,"","",""
DRV:8,256,999,0,"","",""
DRV:9,256,999,0,"","",""
DRV:10,256,999,0,"","",""
DRV:11,256,999,0,"","",""
DRV:12,256,999,0,"","",""
DRV:13,256,999,0,"","",""
DRV:14,256,999,0,"","",""
DRV:15,256,999,0,"","",""
MSG:1011,0,1,"Using LibreDrive mode (v06.3 id=0FA242DD4D0B)","%1","Using LibreDrive mode (v06.3 id=0FA242DD4D0B)"
MSG:3007,0,0,"Using direct disc access mode","Using direct disc access mode"
MSG:5053,1544,1,"This functionality is shareware. You may evaluate it for 30 days after what you would need to purchase an activation key if you like the functionality. Do you want to start evaluation period now?","This functionality is shareware. You may evaluate it for %1 days after what you would need to purchase an activation key if you like the functionality. Do you want to start evaluation period now?","30"
MSG:5050,0,2,"Evaluation version, 30 day(s) out of 30 remaining","Evaluation version, %1 day(s) out of %2 remaining","30","30"
MSG:5085,0,0,"Loaded content hash table, will verify integrity of M2TS files.","Loaded content hash table, will verify integrity of M2TS files."
MSG:3025,0,3,"Title #00001.mpls has length of 12 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00001.mpls","12","120"
MSG:3307,16777216,2,"File 00000.mpls was added as title #0","File %1 was added as title #%2","00000.mpls","0"
MSG:3025,16777216,3,"Title #00017.m2ts has length of 1 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00017.m2ts","1","120"
MSG:3025,16777216,3,"Title #00011.m2ts has length of 1 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00011.m2ts","1","120"
MSG:3025,16777216,3,"Title #00005.m2ts has length of 7 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00005.m2ts","7","120"
MSG:3025,16777216,3,"Title #00004.m2ts has length of 5 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00004.m2ts","5","120"
MSG:3025,16777216,3,"Title #00006.m2ts has length of 9 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00006.m2ts","9","120"
MSG:3025,0,3,"Title #00010.m2ts has length of 39 seconds which is less than minimum title length of 120 seconds and was therefore skipped","Title #%1 has length of %2 seconds which is less than minimum title length of %3 seconds and was therefore skipped","00010.m2ts","39","120"
MSG:5011,0,0,"Operation successfully completed","Operation successfully completed"`
	standardEvents, discInfoEvents := gomakemkv.ParseMakeMkvInfoCommandLogs(strings.NewReader(input))

	standardEventCount := 0
	discInfoEventCount := 0
loop:
	for {
		select {
		case _, ok := <-standardEvents:
			if ok {
				standardEventCount++
			}
		case _, ok := <-discInfoEvents:
			if ok {
				discInfoEventCount++
			}
			break loop
		}
	}
	assert.Equal(t, 32, standardEventCount)
	assert.Equal(t, 0, discInfoEventCount)
}
