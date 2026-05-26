/*
 * MainDataBase.h
 *
 *  Created on: 07.05.2010
 *      Author: dilma
 */

#ifndef MAINDATABASE_H_
#define MAINDATABASE_H_

#include "share.h"

typedef struct {
	ustring name;
	ustring info;
} CommandInfo;

typedef struct {
	int id;
	ustring name;
	ustring info;
	ustring codeName;
	ustring units;
} PackInfo;

typedef struct {
	ustring name;
	ustring info;
	ustring entry;
	ustring ext;
	ustring make;
} ProjectInfo;

typedef struct {
	int id;
	int pid;
	ustring name;
	ustring title;
	ustring info;
	ustring value;
	int type;
	ustring list;
} ParamInfo;

typedef struct {
	int id;
	ustring name;
	ustring cmd;
	ustring path;
	ustring ext;
} CompilerInfo;

typedef struct {
	int compiler;
	bool active;
} PackCompilerInfo;

/*! \class MainDataBase
    \brief access to hiasm.db

    Provides tools for data access and management in hiasm.db
*/

class MainDataBase {
	private:
		bool isOpen;

	public:
		~MainDataBase() { close(); }

		/**
		 * Open hiasm.db
		 */
		void open() { isOpen = false; }
		/**
		 * Close hiasm.db
		 */
		void close() { isOpen = false; }
		/**
		 * API for access application commands (from db table commands)
		 * @return current request id
		 */
		void *begin_read_commands() { return NULL; }
		/**
		 * Read information about command
		 * @param id request (give from begin_read_commands)
		 * @param cmd return information about command (name and descriptor)
		 * @return true if the command successfully read, false - this command is last
		 */
		bool read_commands(void *id, CommandInfo &cmd) { return false; }

		/**
		 * API for access application command menu (from db table menu_commands)
		 * @param name code name from table menus
		 * @return current request id
		 */
		void *begin_read_command_menu(const ustring &name) { return NULL; }
		/**
		 * Read information about menu command
		 * @param id request (give from begin_read_command_menu)
		 * @param cmd return information about command (name and descriptor)
		 * @return true if the command successfully read, false - this command is last
		 */
		inline bool read_command_menu(void *id, CommandInfo &cmd) { return false; }

		/**
		 * API for access commands shortcut (from db table menu_commands)
		 * @param key hardware code
		 * @param shift state of Shift key
		 * @param shift state of Ctrl key
		 * @return current request id
		 */
		void *begin_read_command_by_key(int key, bool shift, bool ctrl) { return NULL; }
		/**
		 * Read information about command
		 * @param id request (give from begin_read_command_by_key)
		 * @param cmd return information about command (name and descriptor)
		 * @return true if the command successfully read, false - this command is last
		 */
		inline bool read_command_by_key(void *id, CommandInfo &cmd) { return false; }

		/**
		 * API for access information about installed packages (from db table packs)
		 * @return current request id
		 */
		void *begin_read_packs() { return NULL; }
		/**
		 * Read information about pack
		 * @param id request (give from begin_read_packs)
		 * @param pack return information about pack
		 * @return true if the pack successfully read, false - this pack is last
		 */
		bool read_packs(void *id, PackInfo &pack) { return false; }

		/**
		 * API for access information about installed projects (from db table projects)
		 * @param name pack codename from table packs
		 * @return current request id
		 */
		void *begin_read_projects(const ustring &name) { return NULL; }
		/**
		 * Read information about project
		 * @param id request (give from begin_read_projects)
		 * @param project return information about project
		 * @return true if the project successfully read, false - this project is last
		 */
		bool read_projects(void *id, ProjectInfo &project) { return false; }

		/**
		 * API for access information about registered settings (from db table params)
		 * @param pid parent id of param
		 * @return current request id
		 */
		void *begin_read_settings(int pid) { return NULL; }
		/**
		 * Read information about params
		 * @param id request (give from begin_read_settings)
		 * @param project return information about param
		 * @return true if the param successfully read, false - this param is last
		 */
		bool read_settings(void *id, ParamInfo &param) { return false; }
		/**
		 * API for access values of settings (from db table params)
		 * @return current request id
		 */
		void *begin_read_settings_value() { return NULL; }
		/**
		 * Read information about value
		 * @param id request (give from begin_read_settings)
		 * @param project return information about param (valid two field: id, value and type)
		 * @return true if the param successfully read, false - this param is last
		 */
		bool read_settings_value(void *id, ParamInfo &param) { return false; }
		/**
		 * Save value to DB
		 * @param id
		 * @rapam value
		 */
		void save_settings_value(int id, const ustring &value) {}

		/**
		 * API for access information about installed compilers (from db table compilers)
		 * @return current request id
		 */
		void *begin_read_compilers() { return NULL; }
		/**
		 * Read information about compiler
		 * @param id request (give from begin_read_projects)
		 * @param compiler return information about project
		 * @return true if the project successfully read, false - this project is last
		 */
		bool read_compilers(void *id, CompilerInfo &compiler) { return false; }
		/**
		 * API for access information about attached compilers to pack (from db table pack_compilers)
		 * @param name codename of pack
		 * @return current request id
		 */
		void *begin_read_pack_compilers(const ustring &name) { return NULL; }
		/**
		 * Read information about compilers
		 * @param id request (give from begin_read_projects)
		 * @param compiler return information about project
		 * @return true if the project successfully read, false - this project is last
		 */
		bool read_pack_compilers(void *id, PackCompilerInfo &compiler) { return false; }

		/**
		 * Close current request
		 */
		void end_read(void *id) {}
};

extern MainDataBase mdb;

#endif /* MAINDATABASE_H_ */
