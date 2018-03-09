#include <SPI.h>
#include "RF24.h"

const int MSG_TIMEOUT = 200;

RF24 radio(7,8);

void readStr(char *buf, uint8_t len)
{
  while(!radio.available());
  radio.read(buf, len);
  buf[len] = '\0'; 
}

void setup()
{
	Serial.begin(9600);
	radio.begin();

	radio.setPALevel(RF24_PA_LOW);

  byte addresses[][6] = {"1Node","2Node"};
	uint8_t write_address = 0x20;
	uint8_t read_address = 0x10;
	radio.openWritingPipe(addresses[0]);
    radio.openReadingPipe(1,addresses[1]);

	radio.startListening();
  Serial.println("Listening for commands");
}

void loop()
{
	if(radio.available())
	{
		uint8_t cmd;
    radio.read(&cmd, sizeof(cmd));
		

		if(cmd == 0x01)
		{
			Serial.print("Recieved name...");

      uint8_t len;
      while(!radio.available());
      radio.read(&len, sizeof(len));
      
      
      while(!radio.available());
      char name[len + 1];
      readStr(name, len);

      Serial.println(name);
		}
		else if(cmd == 0x02)
		{
      Serial.print("Recieved birthday...");
      
      while(!radio.available());
      char birthday[10 + 1];
      readStr(birthday, 10);

      Serial.println(birthday);
		}
		else if(cmd == 0x03)
		{
      Serial.print("Recieved phone...");
      
      while(!radio.available());
      char phone[10 + 1];
      readStr(phone, 10);

      Serial.println(phone);
		}
		else if(cmd == 0x04)
		{
      Serial.print("Recieved status...");

      uint8_t status;
      while(!radio.available());
      radio.read(&status, sizeof(status));
      
      if(status == 0) Serial.println("Uninjured");
      else if(status == 1) Serial.println("Minor injury");
      else if(status == 2) Serial.println("Major injury");
		}
   else if(cmd == 0x05)
   {
     Serial.print("Recieved message...");

      uint8_t len;
      while(!radio.available());
      radio.read(&len, sizeof(len));
      while(!radio.available());
      char message[len + 1];
      readStr(message, len);

      Serial.println(message);
   }
	}
}
