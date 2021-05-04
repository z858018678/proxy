script_path=./scripts
gen:
	@echo ">>>>>>>>>> gen proto <<<<<<<<<<"
	cd $(script_path) && ./gen.sh

update:
	@echo ">>>>>>>>>> update proto <<<<<<<<<<"
	git pull origin master
	cd $(script_path) && ./update_business_operate_type.py && ./gen.sh
	git add .
	git commit -m 'Update AexBusinessOperateType'
	git push origin master
