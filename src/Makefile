CC = g++

CPPFLAGS = -Wall -c -O2 
LDLIBS = 

SOURCES = $(wildcard *.cpp) 
APP = hiasm5

all: $(APP)

clean:
	rm -f $(APP) *.o

$(APP): $(SOURCES:%.cpp=%.o)
	$(LINK.o) $^ $(LDLIBS) -o $(APP)
	strip $@

