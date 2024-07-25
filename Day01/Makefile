EX_00 = DBreader/cmd
EXEC_00 = readDB

EX_01 = DBcompare/cmd
EXEC_01 = compareDB

EX_02 = DBwriter/cmd
EXEC_02 = writeDB

all:
	cd $(EX_00) && go build -o $(EXEC_00)
	cd $(EX_01) && go build -o $(EXEC_01)
	cd $(EX_02) && go build -o $(EXEC_02)

ex00:
	cd $(EX_00) && go build -o $(EXEC_00)

ex01:
	cd $(EX_01) && go build -o $(EXEC_01)

ex02:
	cd $(EX_02) && go build -o $(EXEC_02)


clean00:
	cd $(EX_00) && rm -rf $(EXEC_00)

clean01:
	cd $(EX_01) && rm -rf $(EXEC_01)

clean02:
	cd $(EX_02) && rm -rf $(EXEC_02)

clean:
	cd $(EX_00) && rm -rf $(EXEC_00)
	cd $(EX_01) && rm -rf $(EXEC_01)
	cd $(EX_02) && rm -rf $(EXEC_02)