MODULE_0 = DBreader/cmd
EXECUTABLE_0 = readDB

MODULE_1 = DBcompare/cmd
EXECUTABLE_1 = compareDB

MODULE_2 = DBwriter/cmd
EXECUTABLE_2 = writeDB

ex00:
	cd $(MODULE_0) && go build -o $(EXECUTABLE_0)

ex01:
	cd $(MODULE_1) && go build -o $(EXECUTABLE_1)

ex02:
	cd $(MODULE_2) && go build -o $(EXECUTABLE_2)


clean00:
	cd $(MODULE_0)
	rm -rf $(EXECUTABLE_0)

clean01:
	cd $(MODULE_1)
	rm -rf $(MODULE_1)

clean02:
	cd $(MODULE_2)
	rm -rf $(MODULE_2)