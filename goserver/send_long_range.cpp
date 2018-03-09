#include <cstdlib>
#include <iostream>
#include <RF24/RF24.h>

using namespace std;

RF24 radio(22, 0);

int main(int argc, char **argv)
{
	/*
		long range rf protocol
		0x00					start new person
		0x01, uint, char[...]	send name
		0x02, char[6]			send birthday
		0x03, char[10]			send phone
		0x04, uint				send status
		0x05, uint, char[...]	send message

		0xFF, uint, char[...]	receive live update
	*/
	radio.begin();
	radio.printDetails();
	const uint8_t pipes[][6] = {"1Node","2Node"};
	uint8_t write_address = 0x10;
	uint8_t read_address = 0x20;
	radio.openWritingPipe(pipes[1]);
	radio.openReadingPipe(1, pipes[0]);
	radio.startListening();

	if((argc != 5 && argc != 6)
		|| strlen(argv[2]) != 10		// birthday must be of the form MMDDYY
		|| strlen(argv[3]) != 10	// phone must be of the form AAABBBCCCC
		|| strlen(argv[4]) != 1)		// status must be a single digit
	{
		printf("Usage: send_long_range NAME BIRTHDAY PHONE STATUS [MESSAGE]\n");
		printf("\tNAME\tFull name of the person (first, middle, last)\n");
		printf("\tBIRTHDAY\tMust be in the form of MMDDYY\n");
		printf("\tPHONE\tMust be in the form AAABBBCCCC\n");
		printf("\tSTATUS\tStatus code\n\n");
		printf("Status codes:\n");
		printf("\t0 - The person is OK\n");
		printf("\t1 - The person is not OK (severe injury, etc.)\n");
		printf("\t2 - The person is dead\n\n");
		return 1;
	}

	printf("Name: %s\n", argv[1]);
	printf("Birthday: %s\n", argv[2]);
	printf("Phone: %s\n", argv[3]);
	printf("Status: ");
	if(argv[4][0] - 48 == 0) printf("OK\n");
	else if(argv[4][0] - 48 == 1) printf("Not OK\n");
	else if(argv[4][0] - 48 == 2) printf("Dead\n");
	if(argc == 6) printf("Message: %s\n", argv[5]);
	printf("\n");

	printf("Sending name...");
	uint8_t cmd = 0x01;
	size_t len = strlen(argv[1]);
	radio.stopListening();
	if(!radio.write(&cmd, sizeof(cmd))
		|| !radio.write(&len, sizeof(len))
		|| !radio.write(argv[1], strlen(argv[1])))
	{
		printf("failed\n\n");
		return 1;
	}
	printf("OK\n");

	printf("Sending birthday...");
	cmd = 0x02;
	if(!radio.write(&cmd, sizeof(cmd))
		|| !radio.write(argv[2], 10))
	{
		printf("failed\n\n");
		return 1;
	}
	printf("OK\n");

	printf("Sending phone...");
	cmd = 0x03;
	if(!radio.write(&cmd, sizeof(cmd))
		|| !radio.write(argv[3], 10))
	{
		printf("failed\n\n");
		return 1;
	}
	printf("OK\n");

	printf("Sending status...");
	cmd = 0x04;
	uint8_t status = argv[4][0] - 48;
	if(!radio.write(&cmd, sizeof(cmd))
		|| !radio.write(&status, sizeof(status)))
	{
		printf("failed\n\n");
		return 1;
	}
	printf("OK\n");

	if(argc == 6)
	{
		printf("Sending message...");
		uint8_t cmd = 0x05;
		size_t len = strlen(argv[5]);
		if(!radio.write(&cmd, sizeof(cmd))
			|| !radio.write(&len, sizeof(len))
			|| !radio.write(argv[5], strlen(argv[5])))
		{
			printf("failed\n\n");
			return 1;
		}
		printf("OK\n");
	}
	
	printf("\n");

	return 0;
}
